## Intro
Today, we will learn how to write a test for a simple http server

## Step 1: Simple App
Definition of this app:
- GET /players/{name} should return a number indicating the total number of wins
- POST /players/{name} should record a win for that name, incrementing for every subsequent POST