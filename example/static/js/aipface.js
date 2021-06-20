    var imgData = null;
    var timer = null;

    $(document).ready(function() {

        let $video = $('.video').get(0);
        let $canvas_mask = $('.canvas-mask').get(0);
        let $canvas_img = $('.canvas-screen-short').get(0);
        let $capture = $('.capture');
        let $result_div = $(".result");
        let $error_div = $('.login-error');
        let $checkbox = $('#ispersis');

        let context_mask = $canvas_mask.getContext('2d');
        let context_img = $canvas_img.getContext('2d');
        var mediaStream;
        
        var cw = 200, ch = 200;
        
        var tracker = new tracking.ObjectTracker('face');
        tracker.setInitialScale(4);
        tracker.setStepSize(2);
        tracker.setEdgesDensity(0.1);

        $capture.addClass('noclick');
        
        //tracking
        var tracker = new tracking.ObjectTracker('face');
        tracker.setInitialScale(4);
        tracker.setStepSize(2);
        tracker.setEdgesDensity(0.1);
        
        tracking.track($video, tracker);
        
        tracker.on('track', function(event) {
            context_mask.clearRect(0, 0, cw, ch);
            
            event.data.forEach(function(rect) {
                context_mask.lineWidth = 2;
                context_mask.strokeStyle = '#a64ceb';
                context_mask.strokeRect(rect.x, rect.y, rect.width, rect.height);
                context_mask.font = '11px Helvetica';
                context_mask.fillStyle = "#fff";
                context_mask.fillText('x: ' + rect.x + 'px', rect.x + rect.width + 5, rect.y + 11);
                context_mask.fillText('y: ' + rect.y + 'px', rect.x + rect.width + 5, rect.y + 22);
            });
        });

        //访问用户媒体设备的兼容方法

        function getUserMedia(constraints, success, error) {
            if (navigator.mediaDevices.getUserMedia) {
                //最新的标准API
                navigator.mediaDevices.getUserMedia(constraints).then(success).
                catch (error);
            } else if (navigator.webkitGetUserMedia) {
                //webkit核心浏览器
                navigator.webkitGetUserMedia(constraints, success, error)
            } else if (navigator.mozGetUserMedia) {
                //firfox浏览器
                navigator.mozGetUserMedia(constraints, success, error);
            } else if (navigator.getUserMedia) {
                //旧版API
                navigator.getUserMedia(constraints, success, error);
            }
        }

        function success(stream) {
            //兼容webkit核心浏览器
            let CompatibleURL = window.URL || window.webkitURL;
            //将视频流设置为video元素的源
            mediaStream = stream;

            //video.src = CompatibleURL.createObjectURL(stream);
            $video.srcObject = stream;
            $video.play();

            displayInfo('媒体设备已就绪');
            timerStart();
            $capture.removeClass('noclick');
        }

        function error(error) {
            displayInfo('设备打开失败：' + error.name + ',' + error.message);
        }

        function userMediaClose() {
            mediaStream.getTracks()[0].stop();
        }

        //参考：https://blog.csdn.net/Rachel_ruiqiu/article/details/78614920

        function convertBase64UrlToBlob(urlData) {
            var bytes = window.atob(urlData.split(',')[1]); //去掉url的头，并转换为byte
            //处理异常,将ascii码小于0的转换为大于0
            var ab = new ArrayBuffer(bytes.length);
            var ia = new Uint8Array(ab);
            for (var i = 0; i < bytes.length; i++) {
                ia[i] = bytes.charCodeAt(i);
            }
            return new Blob([ab], {
                type: 'image/jpeg'
            });
        }
        
        function timerStart(){
            timer = setInterval(function(){
                context_img.drawImage($video, 0, 0, cw, ch);
                var type = 'jpeg';
                
                imgData = $canvas_img.toDataURL(type);
            }, 46.666);
        }
        
        function timerStop(){
            window.clearInterval(timer);
        }
        
        try {
            navigator.getUserMedia = (navigator.mediaDevices.getUserMedia || navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia);
    
            if (navigator.getUserMedia) {
                //调用用户媒体设备, 访问摄像头
                getUserMedia({
                    video: {
                        width: 320,
                        height: 320
                    }
                }, success, error);
                displayInfo('媒体设备开启中...');
            } else {
                alert('不支持访问用户媒体');
            }
            
        } catch (e) {
            displayInfo('不支持访问用户媒体');
        }


        //登陆
        $capture.on('click', function() {
            $error_div.remove();
            
            displayInfo('面部匹配中，请稍后...');

            var formDate = new FormData();
            formDate.append("image", convertBase64UrlToBlob(imgData));
            formDate.append("ispersis", $checkbox.is(':checked'));

            $.ajax({
                type: "POST",
                url: "/login",
                data: formDate,
                timeout: 20000,
                contentType: false,
                processData: false,
                success: function(data) {
                    //var obj= $.parseJSON(data);  
                    //login(obj.access, obj.score.toFixed(2));
                    timerStop();
                    if (data != "") {
                        displayInfo('跳转中...');
                        userMediaClose();
                        setTimeout(function(){
                            document.write(data);
                            document.close();
                        }, 100);
                    }
                },
                error: function(data) {
                    console.log("error");
                },
                complete: function(xhr, status) {
                    if (status == 'timeout') {
                        displayInfo("处理超时，请重新匹配！");
                    }
                }
            })
        })

        function displayInfo(info) {
            $result_div.empty();
            $result_div.append('<p>' + info + '</p>')
        }

    })