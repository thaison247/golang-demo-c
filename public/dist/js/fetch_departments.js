var getFormData = (unindexed_array) => {
  var indexed_array = {};

  $.map(unindexed_array, function (n, i) {
    indexed_array[n["name"]] = n["value"];
  });

  return indexed_array;
};

var renderListDepartments = (listDepartments) => {
  $.each(listDepartments, function (i, val) {
    val.department_status = val.active ? "Active" : "Inactive";
    const trElement = `<tr id="department-${val.department_id}">
                            <td id="departmentId">${val.department_id}</td>
                            <td id="departmentName">
                                ${val.department_name}
                            </td>
                            <td id="departmentCode">${val.department_code}</td>
                            <td id="departmentStatus">${
                              val.department_status
                            }</td>
                            <td id="createdAt">${moment(
                              new Date(val.created_at),
                              "LLL"
                            )}</td>
                            <td class="operators" style="text-align: center;">
                                <button id="dep-btn-${
                                  val.department_id
                                }" type="button" class="btn btn-primary detail-btn" data-id="${
      val.department_id
    }">
                                <i class="far fa-edit"></i> Detail
                                </button>
                                <button id="del-dep-btn-${
                                  val.department_id
                                }" type="button" class="btn btn-danger delete-btn" data-id="${
      val.department_id
    }">
                                <i class="far fa-trash-alt"></i> Delete
                                </button>
                            </td>
                        </tr>`;
    $("#departments_table tbody").append(trElement);

    var btnIdSelector = `#dep-btn-${val.department_id}`;
    $(btnIdSelector).click(() => {
      getDepartmentById(val.department_id);
    });

    var btnDelSelector = `#del-dep-btn-${val.department_id}`;
    $(btnDelSelector).click(() => {
      deleteDepartment(val.department_id, val.department_name);
    });
  });

  $("#departments_table").DataTable({
    paging: false,
    lengthChange: false,
    searching: true,
    ordering: true,
    info: true,
    autoWidth: false,
    responsive: true,
    bInfo: false,
    info: false,
  });
};

// get department by id
var getDepartmentById = (departmentId) => {
  var request = $.ajax({
    url: `http://localhost:8080/api/department?departmentid=${departmentId}`,
    method: "GET",
  });

  request.done((msg) => {
    const department = msg.data[0];
    $("#dep-modal").addClass("show");
    $("#dep-modal").css({ display: "block", background: "rgba(0, 0, 0, 0.4)" });

    $("#dep-form #department_id").val(department.department_id);
    $("#dep-form #department_name").val(department.department_name);
    $("#dep-form #department_code").val(department.department_code);
    var active = department.active ? "Active" : "Inactive";
    $("#dep-form #active").val(active);
    var momentDate = moment(department.created_at).format(
      "DD-MM-YYYY hh:mm:ss"
    );
    $("#dep-form #created_at").val(momentDate);
  });

  request.fail(function (jqXHR, textStatus) {
    alert("Request failed: " + textStatus);
  });
};

// get department by id
var deleteDepartment = (departmentId, departmentName) => {
  swal({
    dangerMode: true,
    title: "Are you sure?",
    text: `Delete '${departmentName} - ID: ${departmentId}'`,
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
        delDepartment(departmentId);
        break;
    }
  });
};

var delDepartment = (departmentId) => {
  var request = $.ajax({
    url: `http://localhost:8080/api/department?departmentid=${departmentId}`,
    method: "DELETE",
  });

  request.done((res) => {
    if (res.status == 200) {
      // remove deleted employee row
      $(`#department-${departmentId}`).remove();

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
      text: textStatus,
      icon: "error",
      button: "Close",
    });
  });
};

// update department request
var updateDepReq = (depData) => {
  var updateDepReq = $.ajax({
    url: `http://localhost:8080/api/department?departmentid=${depData.department_id}`,
    method: "PATCH",
    data: JSON.stringify(depData),
    contentType: "application/json",
  });

  updateDepReq.done((res) => {
    if (res.status == 200) {
      $(`#department-${depData.department_id} #departmentName`).text(
        `${depData.department_name}`
      );

      $(`#department-${depData.department_id} #departmentCode`).text(
        `${depData.department_code}`
      );

      var status = depData.active ? "Active" : "Inactive";
      $(`#department-${depData.department_id} #departmentStatus`).text(
        `${status}`
      );

      $("#dep-modal").removeClass("show");
      $("#dep-modal").css({ display: "none", background: "none" });

      swal({
        title: "Successfully!",
        icon: "success",
        button: "OK",
      });
    }
  });

  updateDepReq.fail(function (jqXHR, textStatus) {
    swal({
      title: "Error!",
      text: textStatus,
      icon: "error",
      button: "Close",
    });
  });
};

var addDepReq = (depData) => {
  var addRequest = $.ajax({
    url: `http://localhost:8080/api/department`,
    method: "POST",
    data: JSON.stringify(depData),
    contentType: "application/json",
  });

  addRequest.done((res) => {
    if (res.status == 200) {
      // get inserted employee
      getDepByName(depData.department_name);
    }
  });

  addRequest.fail(function (jqXHR, textStatus) {
    swal({
      title: "Error!",
      text: textStatus,
      icon: "error",
      button: "OK",
    });
  });
};

var getDepByName = (depName) => {
  var getRequest = $.ajax({
    url: `http://localhost:8080/api/department/name?departmentname=${depName}`,
    method: "GET",
  });
  getRequest.done((res) => {
    var newDepartment = res.data[0];
    console.log(newDepartment);
    var status = newDepartment.active ? "Active" : "Inactive";
    var trElement = `<tr id="department-${newDepartment.department_id}">
                            <td id="departmentId">${
                              newDepartment.department_id
                            }</td>
                            <td id="departmentName">
                                ${newDepartment.department_name}
                            </td>
                            <td id="departmentCode">${
                              newDepartment.department_code
                            }</td>
                            <td id="departmentStatus">${status}</td>
                            <td id="createdAt">${moment(
                              new Date(newDepartment.created_at),
                              "LLL"
                            )}</td>
                            <td class="operators" style="text-align: center;">
                                <button id="dep-btn-${
                                  newDepartment.department_id
                                }" type="button" class="btn btn-primary detail-btn" data-id="${
      newDepartment.department_id
    }">
                                <i class="far fa-edit"></i> Detail
                                </button>
                                <button id="del-dep-btn-${
                                  newDepartment.department_id
                                }" type="button" class="btn btn-danger delete-btn" data-id="${
      newDepartment.department_id
    }">
                                <i class="far fa-trash-alt"></i> Delete
                                </button>
                            </td>
                        </tr>`;
    $("#departments_table tbody").append(trElement);

    var btnIdSelector = `#dep-btn-${newDepartment.department_id}`;
    $(btnIdSelector).click(() => {
      getDepartmentById(newDepartment.department_id);
    });

    var btnDelSelector = `#del-dep-btn-${newDepartment.department_id}`;
    $(btnDelSelector).click(() => {
      deleteDepartment(
        newDepartment.department_id,
        newDepartment.department_name
      );
    });

    $("#add-dep-modal").removeClass("show");
    $("#add-dep-modal").css({
      display: "none",
      background: "none",
    });

    swal({
      title: "Successfully!",
      icon: "success",
      button: "OK",
    });
  });

  getRequest.fail((jqXHR, textStatus) => {
    swal({
      title: "Error!",
      text: textStatus,
      icon: "error",
      button: "Close",
    });
  });
};
