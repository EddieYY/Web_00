

$(window).on("load", function() {
		// Animate loader off screen
		$("#preloader").fadeOut("slow");;
	});


/*window.addEventListener("load",  function() {
    $('#preloader').css('display','none');
});*/


$(window).scroll(function(e){
    if($(window).scrollTop()<=0){
        $('.explore,.navbar').addClass('at_top');
    }else{
        $('.explore,.navbar').removeClass('at_top');
    }

});


