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
        $('.select-val').prop('disabled', false);
        $('.input-val').prop('readonly', false);
    }else{
        $('.select-val').prop('disabled', true);
        $('.input-val').prop('readonly', true);
    }
});