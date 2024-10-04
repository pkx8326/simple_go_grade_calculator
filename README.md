# Simple GO Grade Calculator

### Overview
In this project, I programed an exampe of a very simple student grade calculator with the Go Programming Language. The program uses simple if-else blocks to verify the student's grades and output strings from A - F based on an input score (float).

### Program manual
The user will be asked to input a student score in float. The programe takes the score once the user presses Enter or Return button and returns one of the five grades (A, B, C, D, or F) based on the score.

After the calculation, the user will have an option to continue or to stop the program by answering the question with "Y" or "N". The input answer is not case sensitive.

At the step of inputting the score, if the user inputs any number lower than 0 or higher than 100, or if the user input anything non-numerical, the program will notify the user to input a valid score value which is a floating number between 1 - 100. The program will keep notifying the user until a valid score is input.

### Code structure
The code of this program is divided into two main parts: the main program and the grade calculation function. This is to demonstrate the use of functions in the Go Programming Language.

### Program flow
The program flow is shown as the following numbered list:
1. The program asks for a score input (a floating number between 1 - 100)
2. The program validates the input, if the input is of the valid type, the program will then call the grade calculation function and pass the input to that function. If the input is not of the valid type, the program asks the user to give an input again and will keep asking until it gets a valid input.
3. After the grade is caluclated, the program displayed one of the five grades as a string value (A, B, C, D, or F) based on the input score.
4. The program will then ask if the user wants to perform another calculation. The user will have an option to choose an input between "Y" or "N" (not case sensitive). The program will continue from step 1. until it gets "N" or "n" as an input. At the program also verify the user's input at this step and will keep asking until it gets a valid type of an input.
