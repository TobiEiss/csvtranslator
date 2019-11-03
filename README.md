# csvtranslator

## Functions

### PerRow

Gernerates for every row all rows of the template.

Source:
| Column1   | Column2   |
| --------- | --------- |
| Row1 Col1 | Row1 Col2 |
| Row2 Col1 | Row2 Col2 |

Template:
| MyHeader1    | MyHeader2 |
| ------------ | --------- |
| MyStaticVal1 | !Column1  |
| MyStaticVal2 | !Column2  |

Output:
| MyHeader1    | MyHeader2 |
| ------------ | --------- |
| MyStaticVal1 | Row1 Col1 |
| MyStaticVal2 | Row1 Col2 |
| MyStaticVal1 | Row2 Col1 |
| MyStaticVal2 | Row2 Col2 |