
//How Work Carousel 
if ($('.how-work-carousel').length) {
    $('.how-work-carousel').bxSlider({
        pagerCustom: '.pager-one',
        auto: false,
        startSlide: 1,
        infiniteLoop: false,
        easing: 'swing',
        adaptiveHeight: true,
        pause: 5000,
        slideMargin: 50,
        captions: true,
    });
}

// responsive
$('.owl-carousel').owlCarousel({
    loop:true,
    margin:100,
    
    responsiveClass:true,
    responsive:{
        0:{
            items:1,
        },
        600:{
            items:2,
        },
        1000:{
            items:3,
        },
        1000:{
            items:4,
        }
    }
})
