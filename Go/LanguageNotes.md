# Language Notes: Go

These notes go over some takeaways when learning this language. They're a compilation of what I like about the language, what I dislike, and more.

## Elementary takeaways

My likes, dislikes, unsures, and 'things to get used to' from solving all elementary problems.

### Likes

- `switch` cases do not fall through by default. You must declare that you want this behavior for each `case` using the `fallthrough` keyword
- Very lightweight syntax for types. No extra characters.

### Dislikes

- Trailing commas are required
- The syntax for a map type `var x map[string]int`
- Lack of parenthases around conditionals and loops. `for (i := 1; i <= 10; i++)` is invalid, but I wish I had the option to do it... Maybe I am missing something?
- There's no ternary operator. I AM SAD ABOUT THIS!

### Not sure if I like or dislike

- The walrus operator allowing inferred types for declaring and initializing variables. It is useful in some cases, but there is more code consistency and clarity when all variables have their type explicitly stated. It also leads to the engineer not actually knowing the type of the variable, especially when the code "just works."
- `for` loop is _the_ loop. Having no `while` loop makes the code simpler, and I like that. However, I find myself forgetting to increment variables when using `for` in indeterminate loop situations.

### Things to get used to

- Capitalized method signatures. I haven't been using this in any of my code so far, but I probably should go back and clean it up...
- The import system for modules. It's a wide open directory of tools, and I need to get used to the ecosystem to get a feel for tooling. I tried a hashset out in one of the problems I solved but the module I found is not thread safe.
  - I need to understand the `go.mod` file a bit better. I don't think I really grasp its design yet because my use-case was too narrow.

### Things I want to learn

- Can I use streams/mapping instead of writing all these loops? So exhaustive.