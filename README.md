# nmapdojo-plaza
N-Map Dojo :: Plaza - The Main Point 

## Design System
![Plaza-Design](https://user-images.githubusercontent.com/8407237/175818982-9115e82a-81a1-44ef-845c-f30a71777181.jpg)

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

## Todo
* Test
