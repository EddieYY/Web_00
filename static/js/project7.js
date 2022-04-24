

$(window).scroll(function(e){
    if($(window).scrollTop()<=0){
        $('.explore').addClass('at_top');
    }else{
        $('.explore').removeClass('at_top');
    }

});
