$(document).ready(() => {
  var request = $.ajax({
    url: "http://localhost:8080/api/employee/all?limit=10&offset=0",
    method: "GET",
  });

  request.done(function (msg) {
    var listEmployees = msg.data;
    $.each(listEmployees, function (i, val) {
      const trElement = `<tr>
                            <td>${val.employee_id}</td>
                            <td>
                                ${val.full_name}
                            </td>
                            <td>${val.phone_number}</td>
                            <td>${val.department_name}</td>
                            <td>
                              <button type="button" class="btn btn-primary detail-btn" data-id=${val.employee_id} data-toggle="modal" data-target="#detailModal">
                                Detail
                              </button>
                            </td>
                        </tr>`;
      $("#employees_table tbody").append(trElement);
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

$(".detail-btn").click(() => {
  alert("clicked");
});
