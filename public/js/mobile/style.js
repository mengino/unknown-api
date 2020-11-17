;(function(){
    //合集tab切换
    $(".subject .list_tit .list_rexin a").click(function(){
        var index = $(this).index('.subject .list_tit .list_rexin a');
        $(this).addClass("on").siblings("a").removeClass("on");
        $(".theme_list ul").eq(index).addClass('on').siblings().removeClass('on');
    })
    //点击查看
    var flag = true;
    $(".pic-btn").click(function(){
        if( flag ){
            $('.pic-coll .pic-coll-img dd').css('height','auto');
            $(".pic-btn").text("点击收起");
            flag = false;
        }else{
            $('.pic-coll .pic-coll-img dd').css('height','44px');
            $(".pic-btn").text("点击查看更多");
            flag = true;
        }
    });
    //资讯分类
    var flag = true;
    $(".listimgScale1").click(function(){
        if( flag ){
            $(this).addClass("listimgScale2")
            $(".list-zx-b2").css("height","auto")
            flag = false;
        }else{
            $(this).removeClass("listimgScale2")
            $(".list-zx-b2").css("height","75px")
            flag = true;
        }
       
    })
    //分类tab
    /*$(".down-list .tab .tab-wrap a").click(function(){
        $(this).addClass('active').siblings().removeClass('active');
    });*/
    //下载页面的banner
    var mySwiper = new Swiper('.mySwiper',{
        direction : 'horizontal',
        slidesPerView: '1.5',
        spaceBetween: 10,
    }); 
    //点击更多内容
    var flag = true;
    $(".g-info-cont .open-all").click(function(){
        if(flag){
            $(".g-info-main").css("height","auto");
            $(this).addClass("active");
            $(this).html("点击收起" + "<i></i>");
            flag = false;
        }else{
            $(".g-info-main").css("height","340px");
            $(this).removeClass("active");
            $(this).html("点击展开更多" + "<i></i>");
            flag = true;
        }
    });
    var mySwiper2 = new Swiper('.mySwiper2',{
        direction : 'horizontal',
        spaceBetween: 10,
        slidesPerView: '1.5',
    });
    //排行榜tab切换
    $(".rank>.title-list>li").click(function(){
        var index = $(this).index();
        $(this).addClass("on").siblings().removeClass("on");
        $(".rank-list>ul").eq(index).css("display","block").siblings().css("display","none");
    });
    //点击换一换
    var number = 0;
    $(".hot-rank .hot-title .change").click(function(){
        number++
        var len = $(".hot-container ul").length;
        if( number < len){
            // number++;
            $(".hot-container ul").eq(number).removeClass('on').siblings().addClass('on');
        }else{
            number = 0;
            $(".hot-container ul").eq(number).removeClass('on').siblings().addClass('on');
        }
    })
    $(function () {
        $(".page>span .cbtn").css("display","none");
        $(".page>span").click(function(){

            $(this).children("div").show();
        })
        $(document).click(function(){
            $(".page>span .cbtn").css("display","none");
        })
        $(".page>span .cbtn,.page>span").click(function(event){
            event.stopPropagation();
        });
    })
})()