
window.preventDefault = function(e){
	e.preventDefault();
}

$(document).ready(function(){
	var jtlength,jthtml,winwd,leftwz;
	$('#textpimg .bd li').click(function(){
		var temp=$(this).index();
		if(jthtml==null){
			jthtml=$('#textpimg').html();
			jtlength=$('#textpimg .bd li').length;
			winwd=$('.contents').width();
			leftwz=(7.5-0.4*jtlength)/2
		}
		$('.footer').before('<div class="tpqhbox" id="tpqhbox"><div class="tpbox"><div class="tpjzbox"><div class="jtfade" id="jtfade"></div><div class="pimg" id="tcpimg">'+jthtml+'<div class="hd"><ul></ul></div></div></div></div></div><script>$("#tcpimg .hd ul").css("left",'+leftwz+'+"rem")</script>')
		jtfunc(temp)
	})
	function jtfunc(ind){
		TouchSlide({ 
			slideCell: "#tcpimg",
			titCell:".hd ul", //开启自动分页 autoPage:true ，此时设置 titCell 为导航元素包裹层
			mainCell:".bd ul", 
			effect:"leftLoop", 
			interTime:3000,
			autoPage:true,//自动分页
			autoPlay:false, //自动播放
			defaultIndex:ind
		});
		$('#tpqhbox').click(function(){
			$('#tpqhbox').remove()
		})
	}
	
	var f=document.getElementById("slideBox");
    if(f!=null){
      TouchSlide({ 
			slideCell: "#slideBox",
			titCell:".hd ul", //开启自动分页 autoPage:true ，此时设置 titCell 为导航元素包裹层
			mainCell:".bd ul", 
			effect:"leftLoop", 
			interTime:5000,
			autoPage:true,//自动分页
			autoPlay:true //自动播放
		});
      $('.slideBox').removeClass('on');
    }
	// var texts=document.getElementById("textpimg");
 //    if(texts!=null){
 //      TouchSlide({ 
	// 		slideCell: "#textpimg",
	// 		titCell:".hd ul", //开启自动分页 autoPage:true ，此时设置 titCell 为导航元素包裹层
	// 		mainCell:".bd ul", 
	// 		effect:"leftLoop", 
	// 		interTime:3000,
	// 		autoPage:true,//自动分页
	// 		autoPlay:true //自动播放
	// 	});
 //    }
	var d=document.getElementById("rmtj_list");
	var newst=$('#rmtj_list .hd').length;
	var xif=$('#rmtj_list .bd ul').length;
    if(d!=null && newst && xif>1){
      TouchSlide({ 
			slideCell: "#rmtj_list",
			titCell:".hd ul", //开启自动分页 autoPage:true ，此时设置 titCell 为导航元素包裹层
			mainCell:".bd", 
			effect:"leftLoop", 
			interTime:5000,
			autoPage:true,//自动分页
			autoPlay:false //自动播放
		});
    }else{
    	$('#rmtj_list .hd').hide();
    }
	var h=document.getElementById("newsslide");
    if(h!=null){
      TouchSlide({ 
			slideCell: "#newsslide",
			titCell:".hd ul", //开启自动分页 autoPage:true ，此时设置 titCell 为导航元素包裹层
			mainCell:".bd ul", 
			effect:"leftLoop", 
			interTime:5000,
			prevCell:'#btn_prev',
			nextCell:'#btn_next',
			autoPage:true,//自动分页
			autoPlay:true //自动播放
		});
      $('.slideBox').removeClass('on');
    }
	var g=document.getElementById("jt_list");
	var test;
	var teleng;
	var ht;
  //   if(g!=null){
  //     TouchSlide({ 
		// 	slideCell: "#jt_list",
		// 	titCell:".hd ul", //开启自动分页 autoPage:true ，此时设置 titCell 为导航元素包裹层
		// 	mainCell:".bd ul", 
		// 	effect:"leftLoop", 
		// 	interTime:5000,
		// 	prevCell:'#btn_lf',
		// 	nextCell:'#btn_rg',
		// 	autoPage:true,//自动分页
		// 	autoPlay:false //自动播放
		// });
  //     	test=$('#jt_list .hd ul li');
  //     	teleng=test.length;
  //     	ht=parseInt($('.jt_list').height())/2;
  //     	$('#jt_list .hd span i').html(teleng);
  //     	$('.jt_list').css('margin-top',-ht); 
  //   }

    $('.jtbox a').click(function(){
    	// $('.jt_list').show();
    	$('body').addClass('cover-jietu');
    	var ce=$('.jt_list .bd ul li').length;
    	var hth = $(window).height()*0.7;
    	if($('.jt_list .tempWrap').length==0 && ce>1){
	    		TouchSlide({ 
				slideCell: "#jt_list",
				titCell:".hd ul", //开启自动分页 autoPage:true ，此时设置 titCell 为导航元素包裹层
				mainCell:".bd ul", 
				effect:"leftLoop", 
				interTime:5000,
				prevCell:'#btn_lf',
				nextCell:'#btn_rg',
				autoPage:true,//自动分页
				autoPlay:false //自动播放
			});
	    	test=$('#jt_list .hd ul li');
	      	teleng=test.length; 
    	}
    	else{
    		$('#jtbox .hd ul').html('<li class="on"></li>');
    	}
      	$('.jt_list .bd ul li img').css('max-height',hth);
      	ht=parseInt($('.jt_list').height())/2;
      	$('#jt_list .hd span i').html(teleng);
      	// $('.jt_list').css('margin-top',-ht);
    });

    $('.jtfade').click(function(){
    	$('body').removeClass('cover-jietu');
    })

});

 /*$(function(){
	var seatp;
	$('.top_search input').on('focus', function(){
		$('.rmssbox').slideDown(150);
	}).blur(function(){
		$('.rmssbox').slideUp(150);
		seatp=$(this).val();
	});
	$('.menu .a_menu.last').click(function(){
		$('.menu .ul_xl').slideToggle(100);
	})
	$('.bds_mores').click(function(){
		$('body').addClass('cover-share');
		document.addEventListener('touchmove', window.preventDefault, false);
	});
	$('.share_qx').click(function(){
	console.log(1);	
		$('body').removeClass('cover-share');
		document.removeEventListener('touchmove', window.preventDefault, false);
	});
	$('.fade').click(function(){
		$('body').removeClass('cover-share');
		document.removeEventListener('touchmove', window.preventDefault, false);
	})
    $('.wgx .ptxzbox .pgx input').click(function () {
        $('.xzbox.wgx').hide();
        $('.xzbox.gx').show();
    })
    $('.gx .ptxzbox .pgx input').click(function () {
        $('.xzbox.gx').hide();
        $('.xzbox.wgx').show();
    })
	var that;
	$('.cmt_list ul li').each(function(ind,ele){
		$(ele).find('.zan').click(function(){
			$(this).addClass('cli');
			that=$(this);
			setTimeout(function(){remo(that,'cli')},2000);	
		});
		$(ele).find('.zan.ed').click(function(){
			var thas=$(ele).find('.ydzbox');
			$(ele).find('.ydzbox').show().addClass('ydz');
			setTimeout(function(){remo(thas,'ydz'); thas.hide();},2000);	
		});
		$(ele).find('.hf').click(function(){
			$(ele).find('.cmt_blue').toggle();
		})
	});
	$('.xztcbox .close').click(function(){
		$('.fade,.xztcbox').hide();
	});
	$('.xztcbox .tc_list').each(function(ind,ele){
		$(ele).find('p a').click(function(){
			$(ele).find('p a').removeClass('on');
			$(this).addClass('on');
		})
	})
});*/

	// function remo(de,ys){
	// 	de.removeClass(ys);
	// }
// $(document).ready(function(){
// 	var html= document.getElementsByTagName('html')[0];
// 	var zjce=1.44*parseInt(html.style.fontSize);
// 	var ht=$('.game_text .divtxt').height();
// 	var ys;
// 	if($('.game_main').length!=0){
// 		ys=$('.game_main').offset().top;
// 	}
// 	if(ht>zjce){
// 		$('.game_text .divtxt').addClass('on');
// 		$('.article_js .more_btn').show();
// 	}
// 	$('.article_js .more_btn.sq').live('click',function(){
// 		$(this).removeClass('sq');
// 		$(this).addClass('zk');
// 		$(this).html('收起');
// 		$('.article_js .more_div').addClass('on');
// 		$('.game_text .divtxt').removeClass('on');
// 	})
// 	$('.article_js .more_btn.zk').live('click',function(){
// 		$(this).removeClass('zk');
// 		$(this).addClass('sq');
// 		$(this).html('展开');
// 		$('.article_js .more_div').removeClass('on');
// 		$('.game_text .divtxt').addClass('on');
// 		$('body,html').animate({scrollTop:ys},0);
// 	})
// })
// $(function(){
// 	var html= document.getElementsByTagName('html')[0];
// 	var zjce=1.44*parseInt(html.style.fontSize);
// 	var ht=$('.zq_text').height();
// 	var ys;
// 	if($('.zq_infro').length!=0){
// 		ys=$('.zq_infro').offset().top;
// 	}
// 	if(ht>zjce){
// 		$('.zq_text').addClass('on');
// 		$('.zq_infro .more_btn').show();
// 	}
// 	$('.zq_infro .more_btn.sq').live('click',function(){
// 		$(this).removeClass('sq');
// 		$(this).addClass('zk');
// 		$(this).html('收起');
// 		$('.zq_infro .more_div').addClass('on');
// 		$('.zq_text').removeClass('on');
// 	})
// 	$('.zq_infro .more_btn.zk').live('click',function(){
// 		$(this).removeClass('zk');
// 		$(this).addClass('sq');
// 		$(this).html('展开');
// 		$('.zq_infro .more_div').removeClass('on');
// 		$('.zq_text').addClass('on');
// 		$('body,html').animate({scrollTop:ys},0);
// 	})
// })
//
// $(function top(){
// 	//当滚动条的位置处于距顶部100像素以下时，跳转链接出现，否则消失
// 	$(function () {
// 		$(window).scroll(function(){
// 			if ($(window).scrollTop()>300){
// 				$(".back_top").fadeIn(1000);
// 			}
// 			else
// 			{
// 				$(".back_top").fadeOut(1000);
// 			}
// 		});

// 		//当点击跳转链接后，回到页面顶部位置

// 		$(".back_top a").click(function(){
// 			$('body,html').animate({scrollTop:0},500);
// 			return false;
// 		});
// 	});
// });

// $(function(){
// 	$(".select_nav").each(function(ind,ele){
// 		$(this).find(".nav p").click(function(){
// 		var ul=$(ele).find(".new");
// 		if(ul.css("display")=="none"){
// 			ul.slideDown(50);
// 			$('.select_nav').addClass('on');
// 		}else{
// 			ul.slideUp(50);
// 			$('.select_nav').removeClass('on');
// 		}
// 	});
	
// 	$(ele).find(".nav li").click(function(){
// 		var li=$(this).text();
// 		$(ele).find(".nav p em").html(li);
// 		$(ele).find('.new li').removeClass('on');
// 		$(this).addClass('on');
// 		$(ele).find(".new").hide();
// 		$('.select_nav').removeClass('on'); 
// 	});
// 	});
// });


// function shows(c) {
// 	$('#test_' + c + '>li').click(function () {
// 		$(this).siblings().removeClass('on').end().addClass('on');
// 		var i = $(this).index() + 1;
// 		$('.' + c).addClass('hide');
// 		$('.' + c + '_' + i).removeClass('hide');
// 	});
// }
// function show(c) {
// 	$('.kpsj_tab').each(function(ind,ele){
// 		$(ele).find('#test_' + c + '>li').click(function(){
// 			$(this).siblings().removeClass('on').end().addClass('on');
// 			var i = $(this).index() + 1;
// 			$(ele).find('.' + c).addClass('hide');
// 			$(ele).find('.' + c + '_' + i).removeClass('hide');
// 		});
// 	});
// }
// $(function () {
// 	shows('sejg_list');   //推荐栏目
// 	show('kpsj_list');   //推荐栏目

// 	$('.kpsj_box').each(function(ind,ele){
// 		$(ele).find('.kpsj_tit ul li').click(function(){
// 			$(this).addClass('on').siblings().removeClass('on');
// 			var i=$(this).index();
// 			$(ele).find('.kpsj_tab').addClass('hide');
// 			$(ele).find('.kpsj_tab').eq(i).removeClass('hide');
// 		})
// 	})

// });


//友情链接
// $(function(){
// 	var box=document.getElementById("div1"),can=true; 
// 	var html= document.getElementsByTagName('html')[0];
// 	var zjce=0.5*parseInt(html.style.fontSize);
// 	var ht=0.5*parseInt(html.style.fontSize);
// 	var tp=parseInt($('#div1').height());
// 	if(tp>zjce){
// 		$('#div1').css('height',zjce);
// 		box.innerHTML+=box.innerHTML; 
// 		box.onmouseover=function(){can=false}; 
// 		box.onmouseout=function(){can=true}; 
// 		new function (){ 
// 		var stop=box.scrollTop%ht==0&&!can; 
// 		if(!stop)box.scrollTop==parseInt(box.scrollHeight/2)?box.scrollTop=0:box.scrollTop++; 
// 		setTimeout(arguments.callee,box.scrollTop%ht?30:1500); 
// 		}; 	
// 	}	
// })


//  $(function(){
//  	var tp;
//  	$('.sy_list dl .dd_btn a').click(function(){
//  		tp=$(this).parent().find('.ptips');
//  		tp.addClass('ed');
//  		setTimeout(function(){tp.removeClass('ed');},2000);
//  	})

//  	if($('.botblock').length && $('.new_pl').length){
//  		$('body').addClass('xzh-body');
//  	}
//  	if($('.new_pl').length){
//  		$('body').addClass('body-pl')
//  	}

//  })

//  $(function(){
//  	$('#new_pl input').click(function(){
//  		$('body').addClass('cover-cmt');
//  		$('#saytext2').focus();
//  	});
//  	$('#pl_fade .meb,#cancel').click(function(){
//  		$('body').removeClass('cover-cmt');
//  		$('#saytext2').blur();
//  	});

//  	var ht;
//  	if($('.pl_box').length){
//  		ht=$('.pl_box').offset().top;
//  	}
//  	$('.plts').click(function(){
//  		$('body,html').animate({ scrollTop: ht }, 500);
//  	})


//  	var userAgent = navigator.userAgent;
// 	var isios = userAgent.indexOf("iPhone") > -1;
// 	if(isios && $('.new_pl').length!=0){
// 		$('html').addClass('ios');
// 		$(function () {
// 			$('.contents').scroll(function(){
// 				if ($('.contents').scrollTop()>300){
// 					$(".back_top").fadeIn(1000);
// 				}
// 				else
// 				{
// 					$(".back_top").fadeOut(1000);
// 				}
// 			});
// 			$(".back_top a").click(function(){
// 				$('.contents').animate({scrollTop:0},500);
// 				return false;
// 			});
// 		});
// 		var ht;
// 	 	if($('.pl_box').length){
// 	 		ht=$('.pl_box').offset().top;
// 	 	}
// 	 	$('.plts').click(function(){
// 	 		$('.contents').animate({ scrollTop: ht }, 500);
// 	 	})
// 	}

// 	var winht=$(window).height();
// 	$('body').css('height',winht);


//  })

// $(function() {
// 	var userAgent = navigator.userAgent;
// 	var isios = userAgent.indexOf("iPhone") > -1;
// 	if($('.lazy').length){
// 		if($('html').hasClass('ios') && $('.new_pl').length!=0 ){
// 			// $("img.lazy").lazyload({effect : "fadeIn",threshold :150 });
// 			$("img.lazy").lazyload({effect : "fadeIn",container: $(".contents"),threshold :150 });
// 		}
// 		else{
// 			// $("img.lazy").lazyload({effect : "fadeIn"});
// 			$("img.lazy").lazyload({effect : "fadeIn" });
// 		}
// 		$('.lazy').parent().addClass('lazybg');
// 	}
// });
// //更多简介
// $(function() {
//     if ($('#j_app_desc').get(0)) {
//         var descElem = $('#j_app_desc'),
//             descHeight = descElem.height(),
//             descBtn = $('#j_app_desc_btn'),
//             ah = $('#j_app_desc p'),
//             html= document.getElementsByTagName('html')[0],
// 			minHeight=1.8*parseInt(html.style.fontSize);
//             // minHeight = 90;
//         if (descHeight >= minHeight) {
//             // descElem.css('height', minHeight + 'px');
//             descElem.addClass('on');
//             descBtn.attr('rel', 0).show();
//         }
//         descBtn.click(function() {
//             var rel = $(this).attr('rel');
//             if (rel == 0) {
//                 // descElem.css('height', 'auto');
//                 descElem.removeClass('on');
//                 descBtn.html('收起').attr('rel', 1);
//                 $(".more_btn").addClass("bbj");
//             } else {
//                 // descElem.css('height', minHeight + 'px');
//                 descElem.addClass('on');
//                 descBtn.html('更多').attr('rel', 0);
//                 $(".more_btn").removeClass("bbj");
//             }
//         });
//     }
	
	
// 	// 安装弹窗提示
//     $('#viewTap').each(function(){
//     	console.log('viewTap');
//         var $body = $('body');
//         var $html = $('html');
//         var html = '<div class="install-popup" id="installPopup">'+
//             '<div class="bd">'+
//                 '<h3>第三方企业独立游戏安装设置信任方法</h3>'+
//                 '<p><img src="/game/news/statics/assets/skin_img/install/1.jpg" alt=""></p>'+
//                 '<p>设置信任描述文件教程</p>'+
//                 '<p>一、打开游戏弹出下列弹窗，记住文件名</p>'+
//                 '<p><img src="/game/news/statics/assets/skin_img/install/2.jpg" alt=""></p>'+
//                 '<p>二、依次打开设置-通用-描述文件与设备管理</p>'+
//                 '<p><img src="/game/news/statics/assets/skin_img/install/3.jpg" alt=""></p>'+
//                 '<p>三、信任证书</p>'+
//                 '<p><img src="/game/news/statics/assets/skin_img/install/4.jpg" alt=""></p>'+
//                 '<p>如有不懂设置请加客服QQ：97337191询问</p>'+
//             '</div>'+
//             '<div class="close"></div>'+
//         '</div>';
//         $body.append(html);

//         $(this).on('click', function(){
//             $html.addClass('install-popup-cover');
//         })
//         $('#installPopup .close').on('click', function(){
//             $html.removeClass('install-popup-cover');
//         })
//     })
// });

/*20180629*/
// $(document).ready(function(){
// 	var browser = {
//             versions: function () {
//                 var u = navigator.userAgent, app = navigator.appVersion;
//                 return {
//                     ios: !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/),
//                     iPhone: u.indexOf('iPhone') > -1,
//                     iPad: u.indexOf('iPad') > -1,
//                 };
//             }(),
//         }

// var iosurl = $("input[name=iosurl]").val();

// if (browser.versions.ios || browser.versions.iPhone || browser.versions.iPad) {
// 		var ios1zb=$('#ios1zb').val();
// 		var ios2ota=$('#ios2ota').val();
// 		var ios3yc=$('#ios3yc').val();
// 		$(".ptxzbox").removeClass("zyxzbox");
// 		if (iosurl !== "none") {
// 			$(".pxz #ptabtn").addClass("abtn pt");
// 			$(".ptxzbox .pxz").removeClass("none");
// 		} 
// 		if(ios2ota==1){
// 			$('.btn-ios').show();
// 		}else if(ios1zb==1 && ios2ota==0){
// 			$('.ptxzbox .pxz span').html('APP Store官方下载');
// 		}else if(ios3yc==1 && ios2ota==0 && ios1zb==0){
// 			$('.btn-ios').show();
// 		}
// 	}
// });