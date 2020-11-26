$(document).ready(() => {
  var request = $.ajax({
    url: "http://localhost:8080/api/employee/all?limit=10&offset=0",
    method: "GET",
  });

  request.done(function (msg) {
    var listEmployees = msg.data;
    $.each(listEmployees, function (i, val) {
      const trElement = `<tr id="employee-${val.employee_id}">
                            <td>${val.employee_id}</td>
                            <td>
                                ${val.full_name}
                            </td>
                            <td>${val.phone_number}</td>
                            <td>${val.department_name}</td>
                            <td>
                              <button id="emp-btn-${val.employee_id}" type="button" class="btn btn-primary detail-btn" data-id="${val.employee_id}">
                                Detail
                              </button>
                            </td>
                        </tr>`;
      $("#employees_table tbody").append(trElement);

      var btnIdSelector = `#emp-btn-${val.employee_id}`;

      $(btnIdSelector).click(() => {
        getEmployeeById(val.employee_id);
      });
    });

    $("#employees_table").DataTable({
      paging: true,
      lengthChange: false,
      searching: false,
      ordering: true,
      info: true,
      autoWidth: false,
      responsive: true,
    });
  });

  request.fail(function (jqXHR, textStatus) {
    alert("Request failed: " + textStatus);
  });
});

var getEmployeeById = (employeeId) => {
  console.log(`http://localhost:8080/api/employee?employeeid=${employeeId}`);
  var request = $.ajax({
    url: `http://localhost:8080/api/employee?employeeid=${employeeId}`,
    method: "GET",
  });

  request.done(function (msg) {
    const employee = msg.data[0];
    $("#emp-modal").addClass("show");
    $("#emp-modal").css({ display: "block", background: "rgba(0, 0, 0, 0.4)" });

    $("#emp-modal #employee_id").val(employee.employee_id);
    $("#emp-modal #full_name").val(employee.full_name);
    const gender = employee.gender ? "Male" : "Female";
    $("#emp-modal #gender").val(gender);
    $("#emp-modal #email").val(employee.email);
    $("#emp-modal #phone_number").val(employee.phone_number);
    $("#emp-modal #address").val(employee.address);
    $("#emp-modal #job_title").val(employee.job_title);
    $("#emp-modal #department_name").val(employee.department_name);
    $("#emp-modal #effect_from").val(
      moment(new Date(employee.effect_from), "LLL")
    );
  });

  request.fail(function (jqXHR, textStatus) {
    alert("Request failed: " + textStatus);
  });
};

$("#close-btn").click(() => {
  $("#emp-modal").removeClass("show");
  $("#emp-modal").css({ display: "none", background: "none" });
});
$(".close").click(() => {
  $("#emp-modal").removeClass("show");
  $("#emp-modal").css({ display: "none", background: "none" });
});
