(function ($) {
    'use strict';
    $(function () {
        if ($("#performanceLine").length) { 
            const ctx = document.getElementById('performanceLine');
            var graphGradient = document.getElementById("performanceLine").getContext('2d');
            var graphGradient2 = document.getElementById("performanceLine").getContext('2d');
            var saleGradientBg = graphGradient.createLinearGradient(5, 0, 5, 100);
            saleGradientBg.addColorStop(0, 'rgba(26, 115, 232, 0.18)');
            saleGradientBg.addColorStop(1, 'rgba(26, 115, 232, 0.02)');
            var saleGradientBg2 = graphGradient2.createLinearGradient(100, 0, 50, 150);
            saleGradientBg2.addColorStop(0, 'rgba(0, 208, 255, 0.19)');
            saleGradientBg2.addColorStop(1, 'rgba(0, 208, 255, 0.03)');
      
            new Chart(ctx, {
              type: 'line',
              data: {
                labels: ["SUN","sun", "MON", "mon", "TUE","tue", "WED", "wed", "THU", "thu", "FRI", "fri", "SAT"],
                datasets: [{
                  label: 'This week',
                  data: [50, 110, 60, 290, 200, 115, 130, 170, 90, 210, 240, 280, 200],
                  backgroundColor: saleGradientBg,
                  borderColor: [
                      '#1F3BB3',
                  ],
                  borderWidth: 1.5,
                  fill: true, // 3: no fill
                  pointBorderWidth: 1,
                  pointRadius: [4, 4, 4, 4, 4,4, 4, 4, 4, 4,4, 4, 4],
                  pointHoverRadius: [2, 2, 2, 2, 2,2, 2, 2, 2, 2,2, 2, 2],
                  pointBackgroundColor: ['#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)'],
                  pointBorderColor: ['#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff',],
              },{
                label: 'Last week',
                data: [30, 150, 190, 250, 120, 150, 130, 20, 30, 15, 40, 95, 180],
                backgroundColor: saleGradientBg2,
                borderColor: [
                    '#52CDFF',
                ],
                borderWidth: 1.5,
                fill: true, // 3: no fill
                pointBorderWidth: 1,
                pointRadius: [0, 0, 0, 4, 0],
                pointHoverRadius: [0, 0, 0, 2, 0],
                pointBackgroundColor: ['#52CDFF)', '#52CDFF', '#52CDFF', '#52CDFF','#52CDFF)', '#52CDFF', '#52CDFF', '#52CDFF','#52CDFF)', '#52CDFF', '#52CDFF', '#52CDFF','#52CDFF)'],
                  pointBorderColor: ['#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff',],
            }]
              },
              options: {
                responsive: true,
                maintainAspectRatio: false,
                elements: {
                  line: {
                      tension: 0.4,
                  }
                },
              
                scales: {
                  y: {
                    border: {
                      display: false
                    },
                    grid: {
                      display: true,
                      color:"#F0F0F0",
                      drawBorder: false,
                    },
                    ticks: {
                      beginAtZero: false,
                      autoSkip: true,
                      maxTicksLimit: 4,
                      color:"#6B778C",
                      font: {
                        size: 10,
                      }
                    }
                  },
                  x: {
                    border: {
                      display: false
                    },
                    grid: {
                      display: false,
                      drawBorder: false,
                    },
                    ticks: {
                      beginAtZero: false,
                      autoSkip: true,
                      maxTicksLimit: 7,
                      color:"#6B778C",
                      font: {
                        size: 10,
                      }
                    }
                  }
                },
                plugins: {
                  legend: {
                      display: false,
                  }
                }
              },
              plugins: []
            });
          }
    });
    //   iconify(document.querySelector('.my-cool.icon'));
    // });


})(jQuery);