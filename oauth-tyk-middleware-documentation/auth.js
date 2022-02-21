// Create your middleware object
var auth = new TykJS.TykMiddleware.NewMiddleware({});

// Initialise it with your functionality by passing a closure that accepts two objects
// into the NewProcessRequest() function:
auth.NewProcessRequest(function(request, session, spec) {

    // console.log("This middleware does nothing, but will print this to your terminal.")

    var requestObj = {
      "Method": "POST",
      "Headers": {
          "Content-Type": "application/x-www-form-urlencoded"
      },
      "Domain": spec.config_data.TOKEN_SERVER,
      "Resource": spec.config_data.TOKEN_ENDPOINT,
      "FormData": {
          "grant_type": "client_credentials",
          "client_id": spec.config_data.CLIENT_ID,
          "client_secret": spec.config_data.CLIENT_SECRET
      }
  };

  var encodedResponse = TykMakeHttpRequest(JSON.stringify(requestObj));
  var decodedResponse = JSON.parse(encodedResponse);
  var decodedBody = {}

  try {
      decodedBody = JSON.parse(decodedResponse.Body);
  } catch (err) {
      decodedBody = {}
  }
  // request.headers.authorization = 'test';
  request.SetHeaders["authorization"] = "Bearer " + decodedBody.access_token;

  // You MUST return both the request and session metadata
  return auth.ReturnData(request, session.meta_data);
});