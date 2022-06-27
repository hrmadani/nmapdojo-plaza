# nmapdojo-plaza
N-Map Dojo :: Plaza - The Main Point 

## Design System
![System-Design](https://user-images.githubusercontent.com/8407237/176024977-1af202ef-cb35-4370-b124-b0224f1777bc.jpg)


## Api Address
* Show all alive reports : GET /reports
* Add a report : POST /addReport

the Post Json:
{
    "report_id": 1, // report id must generate in 
    "user_id": 2, // the current user id (the 1,2,3,4 user id already exist in mock DB)
    "action": "add", //add - like - dislike
    "report_type" : "" //police - traffic jam - car crash 
}

## MySQL ERD

![ERD](https://user-images.githubusercontent.com/8407237/176046943-22510058-8ce2-4a1e-a61b-b8cacb12a528.jpg)

## Todo
* Test
