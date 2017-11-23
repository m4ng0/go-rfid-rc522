# go-rfid-rc522
A golang library for rc522 rfid reader.

# About
* https://www.youtube.com/watch?v=ZIexZ5r37Zs (PL)
* https://www.youtube.com/watch?v=E_aVF8t5qnw (PL)

# The reason behind this fork
This version of the library introduces a sleep between RFID read operations,
so that your CPU doesn't get eaten too much.
