$(document).ready(function () {
    $('#btn-conectar').click(function () {
        var usuario = {
            username: $("#username").val(),
            password: $("#password").val()
        }

        $.ajax({
            url: 'http://localhost:9000/api/login',
            type: 'POST',
            data: JSON.stringify(usuario),
            success: function () {
                console.log("success");
            },
            error: function (error) {
                console.log(error);
            }
        });
    });
});