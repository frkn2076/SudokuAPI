# SudokuAPI
Sudoku Generator and Validator API developed by Golang and Gin Framework.

* Golang Gin Web Framework https://github.com/gin-gonic/gin
 
 
#### GET &nbsp;&nbsp;&nbsp;&nbsp;&nbsp; .../sudoku/generate/:level &nbsp;&nbsp;(SUDOKU GENERATOR)
#### POST &nbsp;&nbsp; .../sudoku/validate &nbsp;&nbsp;(SUDOKU VALIDATOR)

&nbsp;
<h2 align="center">GENERATOR</h2>
&nbsp;

* You can send :level parameter 1 to 10 to decide the level of sudoku.
* 1 -> Easiest, 10 -> Hardest

<p align="center">
  <img src="https://github.com/frkn2076/SudokuAPI/blob/master/Assets/image2.PNG" width="100%" height="100%">
</p>

&nbsp;
&nbsp;
<h2 align="center">VALIDATOR</h2> 
&nbsp;

* Send completed sudoku(9x9) to see whether it is a valid sudoku.
<p align="center">
  <img src="https://github.com/frkn2076/SudokuAPI/blob/master/Assets/image1.PNG" width="100%" height="100%">
</p>

#### PS: You can edit the default port 3000.
