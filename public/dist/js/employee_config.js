$(document).ready(() => {
  var request = $.ajax({
    url: "http://localhost:8080/api/employee/all?limit=50&offset=0",
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
                            <td id="phoneNumber">${val.email}</td>
                            <td id="departmentName">${val.department_name}</td>
                            <td class="operators" style="text-align: center;">
                                <button id="emp-btn-${val.employee_id}" type="button" class="btn btn-primary detail-btn" data-id="${val.employee_id}">
                                <i class="far fa-edit"></i> Detail
                                </button>
                                <button id="del-emp-btn-${val.employee_id}" type="button" class="btn btn-danger delete-btn" data-id="${val.employee_id}">
                                <i class="far fa-trash-alt"></i> Delete
                                </button>
                            </td>
                        </tr>`;
      $("#employees_table tbody").append(trElement);

      var btnIdSelector = `#emp-btn-${val.employee_id}`;

      $(btnIdSelector).click(() => {
        getEmployeeById(val.employee_id);
      });

      var btnDelSelector = `#del-emp-btn-${val.employee_id}`;

      $(btnDelSelector).click(() => {
        deleteEmployee(val.employee_id, val.employee_name);
      });
    });

    $("#employees_table").DataTable({
      paging: false,
      lengthChange: false,
      searching: true,
      ordering: true,
      info: true,
      autoWidth: false,
      responsive: true,
      info: false,
    });
  });

  request.fail(function (jqXHR, textStatus) {
    alert("Request failed: " + textStatus);
  });
});

var getEmployeeById = (employeeId) => {
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

var deleteEmployee = (employeeId, employeeName) => {
  swal({
    dangerMode: true,
    title: "Are you sure?",
    text: `Delete '${employeeName} - ID: ${employeeId}'`,
    icon: "warning",
    buttons: {
      cancel: "Cancel",
      yes: true,
    },
  }).then((value) => {
    switch (value) {
      case "cancel":
        swal.close();
        break;

      case "yes":
        swal.close();
        delEmployee(employeeId);
        break;
    }
  });
};

var delEmployee = (employeeId) => {
  var request = $.ajax({
    url: `http://localhost:8080/api/employee?employeeid=${employeeId}`,
    method: "DELETE",
  });

  request.done((res) => {
    if (res.status == 200) {
      // remove deleted employee row
      $(`#employee-${employeeId}`).remove();

      swal({
        title: "Deleted successfully!",
        icon: "success",
        button: "OK",
      });
    }
  });

  request.fail((jqXHR, textStatus) => {
    swal({
      title: "Error!",
      icon: "error",
      button: "Close",
    });
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
  $("#add-emp-modal").removeClass("show");
  $("#add-emp-modal").css({ display: "none", background: "none" });
});

$("#add-emp-close-btn").click(() => {
  $("#add-emp-modal").removeClass("show");
  $("#add-emp-modal").css({ display: "none", background: "none" });
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
    swal({
      title: "Error!",
      text: textStatus,
      icon: "error",
      button: "Close",
    });
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
      var getDepReq = $.ajax({
        url: `http://localhost:8080/api/employee?employeeid=${empdepData.employee_id}`,
        method: "GET",
      });

      getDepReq.done((res) => {
        var department = res.data[0];
        if (!department) {
          $(`#employee-${empdepData.employee_id}`).remove();
        } else {
          $(`#employee-${empdepData.employee_id} #departmentName`).text(
            `${department.department_name}`
          );

          $("#add-emp-modal").removeClass("show");
          $("#add-emp-modal").css({ display: "none", background: "none" });

          $("#emp-modal").removeClass("show");
          $("#emp-modal").css({ display: "none", background: "none" });
        }

        swal({
          title: "Successfully!",
          icon: "success",
          button: "OK",
        });
      });
    }
  });

  changeDepReq.fail(function (jqXHR, textStatus) {
    swal({
      title: "Error",
      text: textStatus,
      icon: "error",
      button: "Close",
    });
  });
}

//ADD EMPLOYEE
$("#add-emp-btn").click(() => {
  $("#add-emp-modal form").get(0).reset();
  $("#add-emp-modal form").get(1).reset();
  $("#add-emp-modal").addClass("show");
  $("#add-emp-modal").css({
    display: "block",
    background: "rgba(0, 0, 0, 0.4)",
  });

  var getListDepartmentsReq = $.ajax({
    url: "http://localhost:8080/api/department/all?limit=10&offset=0",
    method: "GET",
  });

  getListDepartmentsReq.done((res) => {
    $("#add-emp_dep-form #input_department_name").empty();
    var listDepartments = res.data;
    $.each(listDepartments, function (i, dep) {
      var option = `<option value="${dep.department_id}">${dep.department_name}</option>`;
      $("#add-emp_dep-form #input_department_name").append(option);
    });
  });

  $("#add-emp_dep-form #input_effect_from").datetimepicker({
    timepicker: false,
    datepicker: true,
    format: "d-m-yy",
  });
});

$("#submit-btn").click(() => {
  var empData = getFormData($("#add-emp-form").serializeArray());
  empData.gender = empData.gender == "Male" ? true : false;

  addEmpReq(empData);
});

function addEmpReq(empData) {
  var addEmpReq = $.ajax({
    url: `http://localhost:8080/api/employee`,
    method: "POST",
    data: JSON.stringify(empData),
    contentType: "application/json",
  });

  addEmpReq.done((res) => {
    if (res.status == 200) {
      // get inserted employee
      getEmpReq(empData);
    }
  });

  addEmpReq.fail(function (jqXHR, textStatus) {
    swal({
      title: "Error!",
      text: textStatus,
      icon: "error",
      button: "OK",
    });
  });
}

function getEmpReq(empData) {
  let req = $.ajax({
    url: `http://localhost:8080/api/employee/email?email=${empData.email}`,
    method: "GET",
  });

  req.done((res) => {
    console.log(res.data);
    const newEmp = res.data[0];
    const trElement = `<tr id="employee-${newEmp.employee_id}">
                            <td id="employeeId">${newEmp.employee_id}</td>
                            <td id="fullName">
                                ${newEmp.full_name}
                            </td>
                            <td id="phoneNumber">${newEmp.phone_number}</td>
                            <td id="phoneNumber">${newEmp.email}</td>
                            <td id="departmentName">${newEmp.department_name}</td>
                            <td style="text-align: center;">
                              <button id="emp-btn-${newEmp.employee_id}" type="button" class="btn btn-primary detail-btn" data-id="${newEmp.employee_id}">
                              <i class="far fa-edit"></i> Detail
                              </button>
                              <button id="del-emp-btn-${newEmp.employee_id}" type="button" class="btn btn-danger delete-btn" data-id="${newEmp.employee_id}">
                              <i class="far fa-trash-alt"></i> Delete
                              </button>
                            </td>
                        </tr>`;
    $("#employees_table tbody").append(trElement);

    var btnIdSelector = `#emp-btn-${newEmp.employee_id}`;

    $(btnIdSelector).click(() => {
      getEmployeeById(newEmp.employee_id);
    });

    var btnDelSelector = `#del-emp-btn-${newEmp.employee_id}`;

    $(btnDelSelector).click(() => {
      swal({
        dangerMode: true,
        title: "Are you sure?",
        text: `Delete '${newEmp.full_name} - ID: ${newEmp.employee_id}'`,
        icon: "warning",
        buttons: {
          cancel: "Cancel",
          yes: true,
        },
      }).then((value) => {
        switch (value) {
          case "cancel":
            swal.close();
            break;

          case "yes":
            swal.close();
            delEmployee(newEmp.employee_id);
            break;
        }
      });
    });

    var empdepData = getFormData($("#add-emp_dep-form").serializeArray());
    empdepData.effect_from =
      moment(empdepData.effect_from, "DD-MM-YYYY").format("YYYY-MM-DD") +
      "T00:00:00Z";
    empdepData.employee_id = Number(newEmp.employee_id);
    empdepData.department_id = Number(empdepData.department_id);
    empdepData.department_name = $(
      "#add-emp_dep-form #input_department_name option:selected"
    ).text();

    changeDepartment(empdepData);
  });

  req.fail(function (jqXHR, textStatus) {
    alert("Request failed: " + textStatus);
    console.log("Request failed: " + textStatus);
  });
}
