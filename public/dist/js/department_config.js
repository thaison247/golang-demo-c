$(document).ready(() => {
  var request = $.ajax({
    url: "http://localhost:8080/api/department/all?limit=10&offset=0",
    method: "GET",
  });

  request.done(function (msg) {
    var listDepartments = msg.data;
    $.each(listDepartments, function (i, val) {
      const trElement = `<tr>
                            <td>${val.department_id}</td>
                            <td>
                                ${val.department_name}
                            </td>
                            <td>${val.department_code}</td>
                            <td>${moment(new Date(val.created_at), "LLL")}</td>
                        </tr>`;
      $("#departments_table tbody").append(trElement);
    });

    $("#departments_table").DataTable({
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
