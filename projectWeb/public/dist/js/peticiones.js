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

                console.log("vas a morrir moe")
                $.ajax({
                    url: 'http://localhost:9000/api/tabla',
                    type: 'GET',
                    success: function (data) {
                        // console.log(data[0])
                        $.each(data, function (key, value) {
                            if (value.tipoTabla == "TABLE") {
                                $(".nav .nav-treeview").append('<li class="nav-item"> ' +
                                    '<a href="#Todocontenido" class="nav-link tableUser" onclick="getTable(`' + value.nombreTabla + '`)">' +
                                    '<i class="far fa-circle nav-icon"></i>' +
                                    '<p>' + value.nombreTabla + '</p></a></li>');
                                // console.log(key + ": " + value.nombreTabla + " --- " + value.tipoTabla);
                            }
                        });
                    },
                    error: function (error) {
                        console.log(error);
                    }
                });

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
            url: 'http://localhost:9000/api/consultas',
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
    $('#btn-desc').click(function () {

    });
    $('.tableUser').click(function () {
        console.log("haga algo");
    });
});

function getTable(nombTabla) {
    console.log(nombTabla);
    var reply = {
        dato: nombTabla
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
}