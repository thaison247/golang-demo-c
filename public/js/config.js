$(document).ready(() => {
  var request = $.ajax({
    url: "http://localhost:8080/api/employee/all?limit=10&offset=0",
    method: "GET",
  });

  request.done(function (msg) {
    var listEmployees = msg.data;
    $.each(listEmployees, function (i, val) {
      $("#list-employees").append(`<li>${val.full_name}</li>`);
    });
  });

  request.fail(function (jqXHR, textStatus) {
    alert("Request failed: " + textStatus);
  });
});
