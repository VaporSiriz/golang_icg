$(function (){
    $("#analyze").click(function(){
        let code = $("#code").val();
        var queryString = $("form[name=codeAnalyze]").serialize() ;
        $.ajax({
            type : 'post',
            url : 'http://127.0.0.1:8080/upload',
            data : queryString,
            dataType : 'multipart/form-data',
            error: function(error){
                console.log(error);
            },
            success : function(json){
                console.log(json)
            }
        });
        return false;
    });
});