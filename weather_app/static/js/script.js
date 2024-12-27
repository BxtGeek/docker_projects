// static/js/script.js
document.addEventListener('DOMContentLoaded', function() {
    const themeToggle = document.getElementById('checkbox');
    const locationInput = document.getElementById('location');
    const suggestionsDiv = document.getElementById('suggestions');
    const weatherInfo = document.getElementById('weather-info');
    
    // Theme switching
    themeToggle.addEventListener('change', function(e) {
        if (e.target.checked) {
            document.body.classList.remove('light-theme');
            document.body.classList.add('dark-theme');
            localStorage.setItem('theme', 'dark');
        } else {
            document.body.classList.remove('dark-theme');
            document.body.classList.add('light-theme');
            localStorage.setItem('theme', 'light');
        }
    });

    // Load saved theme
    const savedTheme = localStorage.getItem('theme') || 'light';
    if (savedTheme === 'dark') {
        document.body.classList.add('dark-theme');
        themeToggle.checked = true;
    }

    // Location search with suggestions
    let debounceTimer;
    locationInput.addEventListener('input', function(e) {
        clearTimeout(debounceTimer);
        const query = e.target.value.trim();
        
        if (query.length < 2) {
            suggestionsDiv.innerHTML = '';
            return;
        }

        debounceTimer = setTimeout(() => {
            fetch(`/suggest/${query}`)
                .then(response => response.json())
                .then(data => {
                    suggestionsDiv.innerHTML = '';
                    data.forEach(location => {
                        const div = document.createElement('div');
                        div.className = 'suggestion-item';
                        div.textContent = `${location.name}`;
                        div.addEventListener('click', () => {
                            locationInput.value = location.name;
                            suggestionsDiv.innerHTML = '';
                            fetchWeather(location.name);
                        });
                        suggestionsDiv.appendChild(div);
                    });
                })
                .catch(error => console.error('Error:', error));
        }, 300);
    });

    // Close suggestions when clicking outside
    document.addEventListener('click', function(e) {
        if (!suggestionsDiv.contains(e.target) && e.target !== locationInput) {
            suggestionsDiv.innerHTML = '';
        }
    });

    function fetchWeather(location) {
        fetch(`/weather/${location}`)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Weather data not found');
                }
                return response.json();
            })
            .then(data => {
                if (data && typeof data.temperature === 'number') {
                    document.getElementById('location-name').textContent = data.location;
                    document.getElementById('temperature').textContent = Math.round(data.temperature);
                    document.getElementById('description').textContent = data.description || 'N/A';
                    document.getElementById('humidity').textContent = data.humidity || 'N/A';
                    document.getElementById('wind-speed').textContent = data.wind_speed || 'N/A';
                    
                    if (data.icon) {
                        document.getElementById('weather-icon').src = `http://openweathermap.org/img/w/${data.icon}.png`;
                    }
                    
                    weatherInfo.classList.remove('hidden');
                } else {
                    throw new Error('Invalid weather data received');
                }
            })
            .catch(error => {
                console.error('Error:', error);
                weatherInfo.classList.add('hidden');
                alert('Could not fetch weather data. Please try again.');
            });
    }
});
