$(document).ready(() => {
  var request = $.ajax({
    url: "http://localhost:8080/api/employee/all?limit=10&offset=0",
    method: "GET",
  });

  request.done(function (msg) {
    var listEmployees = msg.data;
    $.each(listEmployees, function (i, val) {
      const trElement = `<tr id="employee-${val.employee_id}">
                            <td id="employeeId">${val.employee_id}</td>
                            <td id="fullName">
                                ${val.full_name}
                            </td>
                            <td id="phoneNumber">${val.phone_number}</td>
                            <td id="departmentName">${val.department_name}</td>
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

  // $("#employee-nav-link").addClass("active");
});

var getEmployeeById = (employeeId) => {
  console.log(`http://localhost:8080/api/employee?employeeid=${employeeId}`);
  var request = $.ajax({
    url: `http://localhost:8080/api/employee?employeeid=${employeeId}`,
    method: "GET",
  });

  request.done((msg) => {
    var getListDepartmentsReq = $.ajax({
      url: "http://localhost:8080/api/department/all?limit=10&offset=0",
      method: "GET",
    });

    const employee = msg.data[0];
    $("#emp-modal").addClass("show");
    $("#emp-modal").css({ display: "block", background: "rgba(0, 0, 0, 0.4)" });

    $("#emp-form #employee_id").val(employee.employee_id);
    $("#emp_dep-form #employee_id").val(employee.employee_id);
    $("#emp-form #full_name").val(employee.full_name);
    const gender = employee.gender ? "Male" : "Female";
    $("#emp-form #gender").val(gender);
    $("#emp-form #email").val(employee.email);
    $("#emp-form #phone_number").val(employee.phone_number);
    $("#emp-form #address").val(employee.address);
    $("#emp-form #job_title").val(employee.job_title);
    // $("#emp_dep-form #department_name").val(employee.department_name);
    var momentDate = moment(employee.effect_from).format("DD-MM-YYYY");
    $("#emp_dep-form #effect_from").datetimepicker({
      timepicker: false,
      datepicker: true,
      format: "d-m-yy",
      value: momentDate,
    });

    getListDepartmentsReq.done((res) => {
      const listDepartments = res.data;

      $("#emp_dep-form #department_name").empty();

      $.each(listDepartments, function (i, dep) {
        var option =
          dep.department_id == employee.department_id
            ? `<option selected value="${dep.department_id}">${dep.department_name}</option>`
            : `<option value="${dep.department_id}">${dep.department_name}</option>`;
        $("#emp_dep-form #department_name").append(option);
      });
    });

    getListDepartmentsReq.fail(function (jqXHR, textStatus) {
      alert("Request failed: " + textStatus);
    });
  });

  request.fail(function (jqXHR, textStatus) {
    alert("Request failed: " + textStatus);
  });
};

var getFormData = (unindexed_array) => {
  var indexed_array = {};

  $.map(unindexed_array, function (n, i) {
    indexed_array[n["name"]] = n["value"];
  });

  return indexed_array;
};

$("#close-btn").click(() => {
  $("#emp-modal").removeClass("show");
  $("#emp-modal").css({ display: "none", background: "none" });
});

$(".close").click(() => {
  $("#emp-modal").removeClass("show");
  $("#emp-modal").css({ display: "none", background: "none" });
});

$("#save-btn").click(() => {
  var empData = getFormData($("#emp-form").serializeArray());
  empData.gender = empData.gender == "Male" ? true : false;
  empData.employee_id = Number(empData.employee_id);

  updateEmpReq(empData);
});

function updateEmpReq(empData) {
  var updateEmpReq = $.ajax({
    url: `http://localhost:8080/api/employee?employeeid=${empData.employee_id}`,
    method: "PUT",
    data: JSON.stringify(empData),
    contentType: "application/json",
  });

  updateEmpReq.done((res) => {
    if (res.status == 200) {
      console.log(empData);
      $(`#employee-${empData.employee_id} #fullName`).text(
        `${empData.full_name}`
      );
      // $(`#employee-${empData.employee_id} #departmentName`).inner(`${empData.departmentName}`)
      $(`#employee-${empData.employee_id} #phoneNumber`).text(
        `${empData.phone_number}`
      );

      var empdepData = getFormData($("#emp_dep-form").serializeArray());
      empdepData.effect_from =
        moment(empdepData.effect_from, "DD-MM-YYYY").format("YYYY-MM-DD") +
        "T00:00:00Z";
      empdepData.employee_id = Number(empdepData.employee_id);
      empdepData.department_id = Number(empdepData.department_id);
      empdepData.department_name = $(
        "#emp_dep-form #department_name option:selected"
      ).text();

      changeDepartment(empdepData);
    }
  });

  updateEmpReq.fail(function (jqXHR, textStatus) {
    alert("Request failed: " + textStatus);
    console.log("Request failed: " + textStatus);
  });
}

function changeDepartment(empdepData) {
  var changeDepReq = $.ajax({
    url: `http://localhost:8080/api/empdep`,
    method: "POST",
    data: JSON.stringify(empdepData),
    contentType: "application/json",
  });

  changeDepReq.done((res) => {
    if (res.status == 200) {
      console.log(empdepData);
      $(`#employee-${empdepData.employee_id} #departmentName`).text(
        `${empdepData.department_name}`
      );
    }
  });

  changeDepReq.fail(function (jqXHR, textStatus) {
    alert("Request failed: " + textStatus);
    console.log("Request failed: " + textStatus);
  });
}
