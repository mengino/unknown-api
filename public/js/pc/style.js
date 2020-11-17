;(function(){
    //banner轮播
    var mySwiper = new Swiper('.conta-swiper-container',{
        direction : 'horizontal',
        autoplay : 3000,
        pagination : '.swiper-pagination',
        prevButton : '.lb-prev',
        nextButton : '.lb-next',
    });
    //手游排行榜
    $(".bangbox .game_top_app li").mouseover(function(){
        $(this).addClass("on").siblings().removeClass("on");
    });
    //小编推荐
    $(".b-hot-t>ul>li").mouseover(function(){
        $(this).addClass("b-hot-active").siblings().removeClass("b-hot-active");
    }); 
    //本类下载
    $(".con-right .g-white-box>ul li").mouseover(function(){
        $(this).addClass("m-hover").siblings().removeClass("m-hover");
    }); 
    //展开
    var flag = true;
    $(".m-contshow-btn").click(function(){
        if( flag ){
            $(this).text("收起-")
            $(".g-azd-cont").css("height","auto");
            $(".u-content-mh").css("display","none");
            flag = false;
        }else{
            $(this).text("展开+")
            $(".g-azd-cont").css("height","350px");
            $(".u-content-mh").css("display","block");
            flag = true
        }
    });
    //游戏轮播
    var mySwiper2 = new Swiper(".g-previmg-swiper",{
        scrollbar:'.swiper-scrollbar',
        spaceBetween: 20,
        slidesPerView: 'auto',
        freeModeMomentum: false,
        prevButton:'.swiper-button-prev',
        nextButton:'.swiper-button-next',
    });
    $(".g-previmg-swiper").hover(function(){
        $(".u-previmg-btn").show();
        
    },function(){
        hoverTime = window.setTimeout(function(){
            $(".u-previmg-btn").hide();
        },160);
    });
     //初始化
     var scrollHeight = $(window).scrollTop();
     if (scrollHeight >= 200 + 20) {
         $(".g-info-tab").addClass("m-info-float");
     } else {
         $(".g-info-tab").removeClass("m-info-float");
     }
     $(".g-info-tab li").click(function () {
         var index = $(this).index();
         var scrollTo = $(".con-center .g-white-box").eq(index).offset().top;
         $(this).addClass("m-hover").siblings().removeClass("m-hover");
         $("body,html").animate({
             scrollTop: scrollTo - 450
         }, 300);
     });
    //二维码
    $(".m-down-ul").hover(function () {
        $(".hover-img").css('display', 'block');
        $(".hover-img").parents("li").css('height', '170px');
    }, function () {
        $(".hover-img").css('display', 'none');
        $(".hover-img").parents("li").css('height', 'auto');
    });
})();