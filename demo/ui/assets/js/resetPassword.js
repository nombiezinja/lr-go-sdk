const url = location.href;
const params = url.split("?")[1];
const serverUrl = "http://localhost:4000";
let paramsObj = {};

$("#btn-minimal-resetpassword").click(function() {
  if($("#minimal-resetpassword-password").val() !== $("#minimal-resetpassword-confirmpassword").val()) {
    $("#minimal-resetpassword-message").text("Passwords do not match!");
    $("#minimal-resetpassword-message").attr("class", "error-message");
    return;
  }

  let data = {
      "password" : $("#minimal-resetpassword-password").val(),    
      "resettoken" : paramsObj.vtoken
  }
  
  $.ajax({
      method: "PUT",
      data: JSON.stringify(data),
      url: serverUrl + "/login/resetpassword",
      contentType: "application/json",
      error: function(xhr) {
          $("#minimal-resetpassword-message").text(xhr.responseJSON.Description);
          $("#minimal-resetpassword-message").attr("class", "error-message");
      }
  }).done(function() { 
      $("#minimal-resetpassword-message").text("Password reset successful.");
      $("#minimal-resetpassword-message").attr("class", "success-message");
  });
});

if (params) {
  paramsObj = JSON.parse('{"' + decodeURI(params.replace(/&/g, "\",\"").replace(/=/g,"\":\"")) + '"}');

  if (paramsObj.vtype != "reset") {
    window.location.replace("index.html");
  }
} else {
  window.location.replace("index.html");
}
