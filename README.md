# ToastmasterAgenda

**How to Build:**

*Note: This project requires Go version 1.7.4 or higher.*

This package uses oauth2 client for authentication. You need to get service account key from [Google Developer Console](https://console.developers.google.com/project). Place the ``client_secret.json`` to the root of your project.

**Form layout:**

|0| A | B | C |
|-|:---:|:---:|:---:|
|1||Date|Date|
|2|Toastmaster of the Day|FirstName LastName|FirstName LastName|
|3|Joke Master|FirstName LastName|FirstName LastName|
|4|General Evaluator|FirstName LastName|FirstName LastName|
|5|Timer|FirstName LastName|FirstName LastName|
|6|Ah Counter|---|---|
|7|Grammarian|Date|Date|
|8||||
|9|Speaker 1|Nathaniel LastName CC #3 |--N/A--|
|10|Evaluator 1|Chris LastName|--N/A--|
|11|Speaker 2|Betty LastName  CC #6|--N/A--|
|12|Evaluator 2|Carolyn LastName	|--N/A--|
|13|Speaker 3|Ann LastName|--N/A--|
|14|Evaluator 3|FirstName LastName|--N/A--|
|15|Speaker 4|--N/A--||
|16|Evaluator 4|--N/A--||
|17||||
|18|Back Up Speaker(s)|||
|19|||
|20|Table Topics Master|Nick LastName|Chris LastName|
|21|||
|22|Notes||Transition Meeting?|


1. Column A contains the names of the required fields (i.e row 2 is "Toast Master of the Day", row 3 "General Evaluator"). 
2. Columns B and C start out with the Date as the header and the corresponding field corresponding to the rows listed in column A. Example row 2 column B requires the First and Last Name of the "Toast Master of the Day".


**NOTE** :-
i) In order to change the Google Docs sheets, make sure your sheet is '**publicly**' readable with a link.

ii) The spreadsheet identifier which can be found on the shareable link, needs to be added to the file '**Sheets.go**'. For example, if the shareable link is as follows, then the identifier is the highlighted part :-

https://docs.google.com/spreadsheets/d/1_P9K2asfsITSGEAh7PrPLxncSemNnjXHg3_O3q7OW0k/edit?usp=sharing

Thus, the identifier is '**1_P9K2asfsITSGEAh7PrPLxncSemNnjXHg3_O3q7OW0k**' (without the quotes).
