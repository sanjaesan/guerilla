# The Guerilla interpreter
* Tried implementing Pratt parsers(Recursive descent parser), could have 
gone for yacc tool to auto generate parsers but i'll kill the fun 

## Inspiration
Guerilla Interpreter is inspired by @thorsten monkey patched programming language

## Sample syntax
```javascript
let five = 5;
let ten = 10;
// add is an anonymous function
let add = func(x, y) {
x + y;
};
let response = add(five, ten);
```
