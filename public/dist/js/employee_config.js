$(document).ready(() => {
  var request = $.ajax({
    url: "http://localhost:8080/api/employee/all?limit=50&offset=0",
    method: "GET",
  });

  request.done(function (msg) {
    var listEmployees = msg.data;
    $.each(listEmployees, function (i, val) {
      renderEmployeeRow(val);
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
    $("#emp-form #full_name").val(employee.full_name);
    const gender = employee.gender ? "Male" : "Female";
    $("#emp-form #gender").val(gender);
    $("#emp-form #email").val(employee.email);
    $("#emp-form #phone_number").val(employee.phone_number);
    $("#emp-form #address").val(employee.address);
    $("#emp-form #job_title").val(employee.job_title);

    $("#emp_dep-form #employee_id").val(employee.employee_id);
    $("#emp_dep-form-origin #employee_id-origin").val(employee.employee_id);

    var momentDate = moment(employee.effect_from).format("DD-MM-YYYY");
    $("#emp_dep-form #effect_from").datetimepicker({
      timepicker: false,
      datepicker: true,
      format: "d-m-yy",
      value: momentDate,
    });
    $("#emp_dep-form-origin #effect_from-origin").val(momentDate);

    getListDepartmentsReq.done((res) => {
      const listDepartments = res.data;

      $("#emp_dep-form #department_name").empty();
      $("#emp_dep-form-origin #department_name-origin").empty();

      $.each(listDepartments, function (i, dep) {
        var option =
          dep.department_id == employee.department_id
            ? `<option selected value="${dep.department_id}">${dep.department_name}</option>`
            : `<option value="${dep.department_id}">${dep.department_name}</option>`;
        $("#emp_dep-form #department_name").append(option);
        $("#emp_dep-form-origin #department_name-origin").append(option);
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
      $(`#employee-${empData.employee_id} #fullName`).text(
        `${empData.full_name}`
      );
      $(`#employee-${empData.employee_id} #phoneNumber`).text(
        `${empData.phone_number}`
      );
      $(`#employee-${empData.employee_id} #email`).text(`${empData.email}`);

      var empdepData = getFormData($("#emp_dep-form").serializeArray());
      empdepData.effect_from =
        moment(empdepData.effect_from, "DD-MM-YYYY").format("YYYY-MM-DD") +
        "T00:00:00Z";
      empdepData.employee_id = Number(empdepData.employee_id);
      empdepData.department_id = Number(empdepData.department_id);
      empdepData.department_name = $(
        "#emp_dep-form #department_name option:selected"
      ).text();

      var originData = getFormData($("#emp_dep-form-origin").serializeArray());
      originData.effect_from =
        moment(originData.effect_from, "DD-MM-YYYY").format("YYYY-MM-DD") +
        "T00:00:00Z";
      originData.employee_id = Number(originData.employee_id);
      originData.department_id = Number(originData.department_id);
      originData.department_name = $(
        "#emp_dep-form-origin #department_name-origin option:selected"
      ).text();

      if (
        Object.entries(empdepData).toString() ===
        Object.entries(originData).toString()
      ) {
        $("#add-emp-modal").removeClass("show");
        $("#add-emp-modal").css({ display: "none", background: "none" });

        $("#emp-modal").removeClass("show");
        $("#emp-modal").css({ display: "none", background: "none" });

        // renderEmployeeRow(newEmp);
        swal({
          title: "Update employee successfully!",
          icon: "success",
          button: "OK",
        });
      } else {
        changeDepartment(empdepData);
      }
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
  var inputEffectFrom = new Date(empdepData.effect_from);

  $.ajax({
    url: `http://localhost:8080/api/empdep?employeeid=${empdepData.employee_id}`,
    method: "GET",
  }).done((res) => {
    var latestDayStr = new Date(res.data[0].effect_from);
    var latestDay = new Date(latestDayStr);

    if (inputEffectFrom <= latestDay) {
      swal({
        title: "Error",
        text: "Invalid effect_from day",
        icon: "error",
        button: "Close",
      });
    } else {
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
            var newEmp = res.data[0];

            $("#add-emp-modal").removeClass("show");
            $("#add-emp-modal").css({ display: "none", background: "none" });

            $("#emp-modal").removeClass("show");
            $("#emp-modal").css({ display: "none", background: "none" });

            renderEmployeeRow(newEmp);

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
  });
}

//ADD EMPLOYEE
$("#add-emp-btn").click(() => {
  $("#add-emp-modal form").get(0).reset();
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
    $("#add-emp-form #input_department_name").empty();
    var listDepartments = res.data;
    $.each(listDepartments, function (i, dep) {
      var option = `<option value="${dep.department_id}">${dep.department_name}</option>`;
      $("#add-emp-form #input_department_name").append(option);
    });
  });

  $("#add-emp-form #input_effect_from").datetimepicker({
    timepicker: false,
    datepicker: true,
    format: "d-m-yy",
    value: new Date(),
  });
});

$("#submit-btn").click(() => {
  var empData = getFormData($("#add-emp-form").serializeArray());
  empData.gender = empData.gender == "Male" ? true : false;
  empData.effect_from =
    moment(empData.effect_from, "DD-MM-YYYY").format("YYYY-MM-DD") +
    "T00:00:00Z";
  empData.employee_id = Number(empData.employee_id);
  empData.department_id = Number(empData.department_id);
  empData.department_name = $(
    "#add-emp-form #input_department_name option:selected"
  ).text();

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
      let getInsertedEmp = $.ajax({
        url: `http://localhost:8080/api/employee/email?email=${empData.email}`,
        method: "GET",
      });

      getInsertedEmp.done((res) => {
        var resData = res.data[0];
        resData.department_name = empData.department_name;
        $("#add-emp-modal").removeClass("show");
        $("#add-emp-modal").css({ display: "none", background: "none" });
        swal({
          title: "Added employee successfully!",
          icon: "success",
          button: "OK",
        });
        renderEmployeeRow(resData);
      });

      getInsertedEmp.fail((jqXHR, textStatus) => {
        swal({
          title: "Error!",
          text: textStatus,
          icon: "error",
          button: "Close",
        });
      });
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

var renderEmployeeRow = (employee) => {
  var depName = employee.department_name ? employee.department_name : "";
  const trElement = `<tr id="employee-${employee.employee_id}">
  <td id="employeeId">${employee.employee_id}</td>
  <td id="fullName">
      ${employee.full_name}
  </td>
  <td id="phoneNumber">${employee.phone_number}</td>
  <td id="email">${employee.email}</td>
  <td id="departmentName">${depName}</td>
  <td style="text-align: center;">
    <button id="emp-btn-${employee.employee_id}" type="button" class="btn btn-primary detail-btn" data-id="${employee.employee_id}">
    <i class="far fa-edit"></i> Detail
    </button>
    <button id="del-emp-btn-${employee.employee_id}" type="button" class="btn btn-danger delete-btn" data-id="${employee.employee_id}">
    <i class="far fa-trash-alt"></i> Delete
    </button>
  </td>
</tr>`;
  $(`#employee-${employee.employee_id}`).remove();
  $("#employees_table tbody").append(trElement);

  var btnIdSelector = `#emp-btn-${employee.employee_id}`;

  $(btnIdSelector).click(() => {
    getEmployeeById(employee.employee_id);
  });

  var btnDelSelector = `#del-emp-btn-${employee.employee_id}`;

  $(btnDelSelector).click(() => {
    swal({
      dangerMode: true,
      title: "Are you sure?",
      text: `Delete '${employee.full_name} - ID: ${employee.employee_id}'`,
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
          delEmployee(employee.employee_id);
          break;
      }
    });
  });
};
