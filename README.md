# astro_drone_golang

**Instruction**

Follow these step to run project
========================

git clone https://github.com/KaungKhantHtun/astro_drone_golang  

cd astro_drone_golang

make init-dependency  wire gen astro_drone/config


Create Database
=======================

CREATE DATABASE astro_drone_db;


Modify in .env file with your Postgres SQL Server
=====================================

 DB_DSN="host=localhost user=postgres password=postgres dbname=astro_drone_db port=5432"


 Run the Project
=======================

go run .
   

API
=======================

POST => localhost:8080/api/drone/

GET => localhost:8080/api/drone/

GET => localhost:8080/api/drone/1994  


Payload 
======================= 

[
   {
      "drone_id":29843,
      "instructions":[
         {
            "action":"R",
            "times":1
         },
         {
            "action":"F",
            "times":2
         },
         {
            "action":"R",
            "times":2
         },
         {
            "action":"F",
            "times":4
         },
         {
            "action":"R",
            "times":1
         },
         {
            "action":"L",
            "times":1
         }
      ]
   },
   {
      "drone_id":1994,
      "instructions":[
         {
            "action":"F",
            "times":1
         },
         {
            "action":"F",
            "times":1
         },
         {
            "action":"R",
            "times":2
         },
         {
            "action":"F",
            "times":3
         },
         {
            "action":"R",
            "times":1
         },
         {
            "action":"L",
            "times":1
         }
      ]
   }
]  




