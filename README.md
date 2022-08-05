## Java Lines of Code
Create a function ```CountLinesOfJavaCode``` that takes a Java program as a string argument and returns the actual number of lines of code.

## Description
Counting the number of lines in a text file is easy. The rules for a line of java code however are a bit more complicated. For this exercise we only want to count a line as a line of code, if it contains something other than whitespace or a comment.

Keep in mind that Java comments are either precedented by ```//``` or enclosed with ```/* <code> */```. The suffix ```//``` only works for the text that follows on the same line, while ```/* <code> */``` can also enclose blocks that span multiple lines.

## Examples:

``` java
  -   // This file contains 3 lines of code
  1   public interface Dave {
  -     /**
  -      * count the number of lines in a file
  -      */
  2     int countLinesOfJavaCode(String theProgram); // not the real signature!
  3   }
```

``` java
  -   /*****
  -    * This is a test program with 5 lines of code
  -    //*****//***/// Slightly pathological comment ending...
  -
  1  public class Hello {
  2      public static final void main(String [] args) { // gotta love Java
  -        // Say hello
  3        System./*wait*/out./*for*/println/*it*/("Hello/*");
  4      }
  -
  5  }
```