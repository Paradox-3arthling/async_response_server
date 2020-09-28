# async_response_server
> Sorry for the long name :D
- This is a library that makes a server  that runs in the background. The caller can then use it for callback based API's*,
  which then transfers the information to the caller using channels.
  
 - * callback based API's - where endpioint does not know when a request will be completed(It just validates your data)
   then when completed it will send the response to the callback URL you set.
