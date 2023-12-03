# Alle_
Assignment Project

Problem :
Build a prototype for a chatbot that allows users to save and retrieve images. Also, enable the chatbot to have a discussion with the user using off-the-shelf language model solutions. 

Use golang for your server.
Publish your code on github.

Expectations
- Figure out which are the best off-the-shelf solutions
- Build a functional prototype, UI is good to have. If not, should be demo-ready through APIs
- Integrate off-the-shelf language models for conversational capability

Solution is this github repo: 
Here below are the three curls for the project: 

curl --location 'http://localhost:8080/upload' \
--form 'image=@"/Users/gauravt/Downloads/wallpapers/wp2570425.jpg"'

curl --location 'http://localhost:8080/image/wp2570425.jpg'

curl --location 'http://localhost:8080/chat' \
--header 'Content-Type: application/json' \
--data '{
    "message": "hi"
}'


const openaiAPIKey = "YOUR_OPEN_AI_KEY"

Replace this YOUR_OPEN_AI_KEY with valid openAi key to run the chat system.

