function login() {
    var usuario = {
        username: $("#username").val(),
        password: $("#password").val()
    }

    //$('#target').html('sending..');

    $.ajax({
        url: 'http://localhost:8080/api/login',
        data: JSON.stringify(usuario),
        type: 'POST',
        dataType: 'json',
        contentType: 'application/json',
        success: function (data) {
            //$('#target').html(data.msg);
            alert(data)
        },
        error: function (error) {
            console.log(error);
        }
    });

}