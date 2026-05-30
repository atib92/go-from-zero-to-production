# Notes

## Structs Are Values

Struct assignment copies data.

Example:

user2 := user1

creates a new struct.

---

## No Classes

Go intentionally avoids classes.

Instead:

- structs hold data
- methods provide behavior

---

## Composition Over Inheritance

Go favors:

Employee contains User

instead of:

Employee extends User

This is a major Go philosophy.

# Methods

Methods are functions attached to a type.

Example:

func (u User) Greeting() string

The receiver appears between:
- func
- method name

Methods allow behavior to be associated with data.

---

## Value Receivers

func (u User)

The struct is copied.

Mutations affect only the copy.

Useful for:
- read-only methods
- small structs

---

## Pointer Receivers

func (u *User)

The method operates on the original struct.

Useful for:
- mutations
- large structs
- avoiding copies

---

## Automatic Dereferencing

Go automatically converts:

user.Birthday()

into:

(&user).Birthday()

when needed.

This makes method calls ergonomic.

---

## Important Insight

Methods are the foundation for interfaces.

Interfaces are satisfied by methods, not by fields.


## Methods Feel Similar to OOP Methods

Methods in Go feel similar to methods on classes in traditional OOP languages such as Python.

Example:

Python:

```python
class User:
    def greeting(self):
        return f"Hello {self.name}"
```

Go:

```go
type User struct {
    Name string
}

func (u User) Greeting() string {
    return "Hello " + u.Name
}
```

In both cases, behavior is associated with data.

The major difference is that Go separates:
- data definition (`struct`)
- behavior definition (`method`)

instead of nesting behavior inside a class definition.

---

## Methods Look Like Functions

A Go method signature looks much closer to a function than a traditional OOP method.

Example:

```go
func (u User) Greeting() string
```

This reinforces the idea that methods are essentially ordinary functions with a receiver attached.
The language keeps functions and methods conceptually very close to each other.

---

## The Receiver Attaches Behavior to a Type

The receiver:

```go
(u User)
```

acts as the mechanism that associates a function with a type.
Without a receiver:

```go
func Greeting(u User) string
```

the code is just a normal function.
With a receiver:

```go
func (u User) Greeting() string
```

the function becomes part of the behavior of `User`.
A useful mental model is:

```text
receiver = attach function to type
```

---

## Methods Feel Like Functions Receiving the Struct

A method can be mentally viewed as a regular function where the struct instance is passed as an argument.

Example:

```go
user.Greeting()
```

can be thought of conceptually as:

```go
Greeting(user)
```

This mental model helps explain:
- value receivers
- pointer receivers
- method calls

---

## Pointer Receivers Are Ergonomic

Go automatically dereferences struct pointers.

Example:

```go
func (u *User) Birthday() {
    u.Age++
}
```

Inside the method:

```go
u.Age
```

works directly.

There is no need for:

```go
(*u).Age
```

even though that is conceptually what is happening.

Similarly:

```go
user.Birthday()
```

works even though `Birthday()` is defined on `*User`.

Go automatically takes the address when possible.
This makes pointer-based code significantly cleaner (less error prone) than languages such as C.

---

## Go's Philosophy Feels Different from Traditional OOP

So far, Go appears to favor:

```text
Data + Behavior + Composition
```

rather than:

```text
Classes + Inheritance + Hierarchies
```

Structs hold data.

Methods provide behavior.

Composition is used to build larger abstractions.

This feels simpler and more explicit than traditional inheritance-heavy object-oriented designs.