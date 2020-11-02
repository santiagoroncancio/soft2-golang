function login() {
    $.ajax({
        // url: 'http://localhost:8080/api/login',
        url: 'http://localhost:8080/api/notes',
        type: 'GET'
    }).done(function(data){
    alert(data);
    });
}
// function login() {
//     var usuario = {
//         "title": "Titulo 1",
//         "description": "Description"
//     }

//     $.ajax({
//         // url: 'http://localhost:8080/api/login',
//         url: 'http://localhost:8080/api/notes',
//         type: 'POST',
//         contentType: 'application/json;charset = utf-8',
//         data: JSON.stringify(usuario),
//         success: function (data) {
//             console.log(data);
//         },
//         error: function(error){
//             console.log(error);
//         }
//     });
// }