$('.con').click(function(){
    alert("POSTします");
})

$('.del').click(function(){
    var result = confirm("本当に削除しますか");
    if (!result) {
        return false;
    }
})


