Download the youtube package
- go get google.golang.org/api/youtube/v3

1. Obtain a YouTube Data API Key: You need to obtain an API key from the Google Cloud Console to access the YouTube Data API. You can follow the instructions here to create a project, enable the YouTube Data API, and generate an API key.

2. Replace "YOUR_API_KEY": Replace "YOUR_API_KEY" in the code with the API key you obtained from the Google Cloud Console.

3. Replace "VIDEO_ID": Replace "VIDEO_ID" with the actual ID of the YouTube video you want to download. You can find the video ID in the URL of the video after the v= parameter.

The script will use the API key provided to make a request to YouTube's servers and retrieve the video title via the key to download the video. It will then save the video content to a file with the same name as the title as an `.mp4` extension. 

Reference
- Youtube Data API Reference (https://developers.google.com/youtube/v3/docs/?apix=true)