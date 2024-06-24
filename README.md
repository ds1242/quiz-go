# Go Command Line Quiz
This command line program has a default csv file that it parses the questions and answers or a path to a similar csv file can be given to parse. It then displays the questions to the user and checks if the provided answer is correct.


## Enhancements
- add strings package to trim input to ensure whitespace isn't counted against users
- add timer
    - default is 30 seconds
    - time should be adjustable with a flag
- require enter key to start quiz and timer

### Assumptions
This currently assumes that a csv with two columns is provided for questions and answers.  Currently will not check the first row for titles so the first row must contain quiz information as well.

