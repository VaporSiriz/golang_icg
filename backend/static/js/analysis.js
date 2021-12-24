
$("#analysis").click(function() {
    let code = $("#code").val();
    $.ajax({
        method: "GET",
        url: 'http://127.0.0.1:8080/upload',
        contentType: 'text/plain; charset=utf-8',
        dataType: 'multipart/form-data; charset=utf-8',
        data: {'code': code},
        error : function(error) {
            console.log(error);
        },
        success : function(data) {
            console.log(data);
        },
        complete : function(data) {
            // 왜 error가 뜨는지는 모르겠지만 일단 임시로:
            let result = $("#result");
            result.text(data.responseText);
        }
    }).done(function( msg ) {
        let result = $("#result");
        console.log( msg );
        result.text("test");
    })
});
