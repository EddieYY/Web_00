

$(window).scroll(function(e){
    if($(window).scrollTop()<=0){
        $('.explore,.navbar').addClass('at_top');
    }else{
        $('.explore,.navbar').removeClass('at_top');
    }

});
