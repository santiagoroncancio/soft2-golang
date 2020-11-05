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
            success: function (data) {
                alert(data.dato);
            },
            error: function (error) {
                console.log(error);
            }
        });
    });

    // Error en el boton ejecutar por la cantidad de parametros en la consulta
    $('#btn-ejecutar').click(function () {
        var reply = {
            dato: $("#text-area").val()
        }

        $.ajax({
            url: 'http://localhost:9000/api/consulta',
            type: 'POST',
            data: JSON.stringify(reply),
            success: function (data) {
                alert(data.dato);
            },
            error: function (error) {
                console.log(error);
            }
        });
    });

    $('#btn-limpiar').click(function () {
        $("#text-area").val("");
    });
    $('#btn-desc').click(function(){
        console.log("vas a morrir moe")
        $.ajax({
            url: 'http://localhost:9000/api/tabla',
            type: 'GET',
            success: function (data) {
                alert(data.dato);
            },
            error: function (error) {
                console.log(error);
            }
        });
    });
});