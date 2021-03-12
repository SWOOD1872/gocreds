# gocreds

## Prompt

Prompts for input which is not sensitive e.g. username

```go
u, err := Prompt("Enter username")
if err != nil {
    ...
}
```

&nbsp;

## PromptSecret

Prompts user for input which is sensitive e.g. password

```go
u, err := Prompt("Enter password")
if err != nil {
    ...
}
```