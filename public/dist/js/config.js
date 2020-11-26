$(document).ready(() => {
  var pathname = window.location.pathname;
  if (pathname == "/employee") {
    $("#employee-nav-link").addClass("active");
    $("#department-nav-link").removeClass("active");
  } else {
    $("#department-nav-link").addClass("active");
    $("#employee-nav-link").removeClass("active");
  }
});
