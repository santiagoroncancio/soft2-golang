$(document).ready(function () {
    $('#btn-conectar').click(function () {
        var usuario = {
            "title": "Titulo 1",
            "description": "Description"
        }

        $.ajax({
            // url: 'http://localhost:8080/api/login',
            url: 'http://localhost:9000/api/notes',
            type: 'POST',
            data: JSON.stringify(usuario),
            success: function () {
                console.log("Funciona");
            },
            error: function (error) {
                console.log(error);
            }
        });
    });
});