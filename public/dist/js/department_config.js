$(document).ready(() => {
  var request = $.ajax({
    url: "http://localhost:8080/api/department/all?limit=15&offset=0",
    method: "GET",
  });

  request.done(function (msg) {
    var listDepartments = msg.data;
    renderListDepartments(listDepartments);
  });

  request.fail(function (jqXHR, textStatus) {
    alert("Request failed: " + textStatus);
  });
});

$("#close-btn").click(() => {
  $("#dep-modal").removeClass("show");
  $("#dep-modal").css({ display: "none", background: "none" });
});

$(".close").click(() => {
  $("#dep-modal").removeClass("show");
  $("#dep-modal").css({ display: "none", background: "none" });
  $("#add-dep-modal").removeClass("show");
  $("#add-dep-modal").css({ display: "none", background: "none" });
});

$("#add-dep-close-btn").click(() => {
  $("#add-dep-modal").removeClass("show");
  $("#add-dep-modal").css({ display: "none", background: "none" });
});

$("#save-btn").click(() => {
  var depData = getFormData($("#dep-form").serializeArray());
  depData.department_id = Number(depData.department_id);

  updateDepReq(depData);
});

$("#add-dep-btn").click(() => {
  $("#add-dep-modal form").get(0).reset();
  $("#add-dep-modal").addClass("show");
  $("#add-dep-modal").css({
    display: "block",
    background: "rgba(0, 0, 0, 0.4)",
  });
});

$("#submit-btn").click(() => {
  var depData = getFormData($("#add-dep-form").serializeArray());

  addDepReq(depData);
});
