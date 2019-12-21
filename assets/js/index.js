$('.con').click(function(){
    alert("POSTします");
})

$('.del').click(function(){
    var result = confirm("本当に削除しますか");
    if (!result) {
        return false;
    }
})



$(function(){
    var btnSize = $('button').length;
    if(btnSize > 0){
        $('.input-val').prop('disabled', false);
    }else{
        $('.input-val').prop('disabled', true);
    }
});