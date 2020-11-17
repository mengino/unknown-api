
/*搜索框*/
// if (navigator.userAgent.indexOf('Mac') >= 0 && location.pathname.indexOf("/news/") < 0 && location.pathname.indexOf("search") < 0) {
// 	$("#top .sbox").html('<input class="s_text">请输入游戏名</input><a href="http://m.yiwan.com/search?wd=请输入游戏名"  class="s_btn"></a>') //IOS
// } else {
// 	$("#top .sbox").html('<a href="###" class="s_text">请输入游戏名</a><a href="http://m.yiwan.com/search?wd=请输入游戏名"  class="s_btn"></a>') //安卓
// }
(function (doc, win) {
    var docEl = doc.documentElement,
        resizeEvt = 'orientationchange' in window ? 'orientationchange' : 'resize',
        recalc = function () {
            var clientWidth = docEl.clientWidth;
            if (!clientWidth) return;
            docEl.style.fontSize = 20 * (clientWidth / 375) + 'px';
        };
    if (!doc.addEventListener) return;
    win.addEventListener(resizeEvt, recalc, false);
    doc.addEventListener('DOMContentLoaded', recalc, false);
})(document, window);

$(function () {
    //头部可滑动
    var TjZz = function () {
        if ($('.tj-header').length > 0) {
            var swiper = new Swiper('.tj-header', {
                slidesPerView: 5.7,
                pagination: {
                    el: '.swiper-pagination',
                    clickable: true,
                },
            });
/*            $(".tj-header .swiper-slide").click(function () {
                var index_ = $(this).index();
                $(".swiper-wrapper .swiper-slide").eq(index_).addClass("active").siblings().removeClass("active");
                swiper.slideTo(index_,500,false);
            })*/
        }
    }
    TjZz();
})
