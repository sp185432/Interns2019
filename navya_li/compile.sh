g++ -o realCalculatorString realCalculatorString.cpp ; echo COMPILATION SUCCESSFUL ; mkdir debug ; mv realCalculatorString debug ; echo BUILD SUCCESSFUL ; cd debug ; echo  Build Sucessful, Prerequisites: Install DevC++ compiler to execute the program >>log.log ; cd .. ; (echo cd debug ; echo ./realCalculatorString ; echo cd .. ; echo echo Execution Successfully Completed >>runLOG.log) >run.sh ; echo Build Successful, execute the "run.sh" file >>compileLOG.log
