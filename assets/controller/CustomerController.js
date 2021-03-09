loadAllCustomers()


$('#btnClr').click(
    function (){
        console.log("hidsfasfsadfas")
        $("#cus_id").val("");
        $("#cus_name").val("");
        $("#cus_nic").val("");
        $("#cus_contact_number").val("");
        $("#cus_address").val("");
        $('#dispimg').attr('src', '');
    }
);
//loadCustomers to the table
function loadAllCustomers() {
    $('#tblCustomer').empty();
    $.ajax({
        method: "GET",
        url: "http://localhost:8000/api/customer",
        success: function (res) {
            let data = res.data;
            //console.log(data)
            for (var i in res){
                let id = res[i].id;
                let fullname = res[i].full_name;
                let nic = res[i].nic;
                let contact = res[i].contact_number;
                let address = res[i].address;
                let filepath = res[i].img_file_path;
                var img = '<img width="30" src="' + filepath + '"/>';
                var row=`<tr> <td>${id}</td> <td>${fullname}</td><td>${nic}</td><td>${contact}</td><td>${address}</td><td>${img}</td></tr>`;
                $('#tblCustomer').append(row);
            }
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    })
}

//Search Customer
$("#btnSrch").click(function () {
    let id = $("#cus_id").val();
    $.ajax({
        method: "GET",
        url: "http://localhost:8000/api/customer/" + id,
        success: function (res) {
            console.log(res);
            let c = res.data;
            $("#cus_name").val(res.full_name);
            $("#cus_nic").val(res.nic);
            $("#cus_contact_number").val(res.contact_number);
            $("#cus_address").val(res.address);
            $('#dispimg').attr('src',res.img_file_path);
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    });
});

//Add Customer
$("#btnAdd").click(function (event) {
    let id = $("#cus_id").val();
    let fullname = $("#cus_name").val();
    let nic = $("#cus_nic").val();
    let contact = $("#cus_contact_number").val();
    let address = $("#cus_address").val();
    //stop submit the form, we will post it manually.
    event.preventDefault();

    // Get form
    var form = $('form').get(0);

    // Create an FormData object
    var data = new FormData(form);
    //var data = new FormData(document.getElementById("btnImgUp"));
    console.log("datas");

    // // If you want to add an extra field for the FormData
    data.append("cus_id",id);
    data.append("cus_name",fullname);
    data.append("cus_nic",nic);
    data.append("cus_contact_number",contact);
    data.append("cus_address",address);

    // disabled the submit button
    $("#btnImgUp").prop("disabled", true);

    $.ajax({
        type: "POST",
        enctype: 'multipart/form-data',
        url: "http://localhost:8000/api/customer",
        data: data,
        processData: false,
        contentType: false,
        cache: false,
        timeout: 600000,
        success: function (data) {
            loadAllCustomers()
            console.log("SUCCESS : ", data);
            $("#btnImgUp").prop("disabled", false);
        },
        error: function (e) {
            console.log("ERROR : ", e);
            $("#btnImgUp").prop("disabled", false);

        }
    });

});


//Update Customer
$("#btnUpdt").click(function (event) {
    let id = $("#cus_id").val();
    let fullname = $("#cus_name").val();
    let nic = $("#cus_nic").val();
    let contact = $("#cus_contact_number").val();
    let address = $("#cus_address").val();
    //stop submit the form, we will post it manually.
    event.preventDefault();

    // Get form
    var form = $('form').get(0);

    // Create an FormData object
    var data = new FormData(form);
    //var data = new FormData(document.getElementById("btnImgUp"));

    // If you want to add an extra field for the FormData
    data.append("cus_id",id);
    data.append("cus_name",fullname);
    data.append("cus_nic",nic);
    data.append("cus_contact_number",contact);
    data.append("cus_address",address);

    // disabled the submit button
    $("#btnImgUp").prop("disabled", true);

    $.ajax({
        type: "PUT",
        enctype: 'multipart/form-data',
        url: "http://localhost:8000/api/customer/"+id,
        data: data,
        processData: false,
        contentType: false,
        cache: false,
        timeout: 600000,
        success: function (data) {
            loadAllCustomers()
            console.log("SUCCESS : ", data);
            $("#btnImgUp").prop("disabled", false);
        },
        error: function (e) {
            console.log("ERROR : ", e);
            $("#btnImgUp").prop("disabled", false);

        }
    });

});



//Delete Customer
$("#btnDlt").click(function (){
    let id=$("#cus_id").val();
    $.ajax({
        method:"DELETE",
        url:"http://localhost:8000/api/customer/"+ id,
        success:function (res){
            alert("the customer is removed");
            loadAllCustomers()
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    });
});

function readURL(input) {
    if (input.files && input.files[0]) {
        var reader = new FileReader();

        reader.onload = function(e) {
            $('#dispimg').attr('src', e.target.result);
        }

        reader.readAsDataURL(input.files[0]); // convert to base64 string
    }
}

$("#btnImgUp").change(function() {
    readURL(this);
});