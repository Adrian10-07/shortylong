const API_BASE_URL = 'http://localhost:8080';

const personaForm = document.getElementById('personaForm');
const recentPersonasContainer = document.getElementById('recentPersonas');
const maleCountElement = document.getElementById('maleCount');
const femaleCountElement = document.getElementById('femaleCount');

let recentPersonas = [];

async function registerPersona(personaData) {
    try {
        const response = await fetch(`${API_BASE_URL}/personas`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(personaData)
        });
        
        if (!response.ok) {
            throw new Error('Error al registrar persona');
        }
        
        const data = await response.json();
        console.log('Persona registrada:', data);
        
        recentPersonas.unshift(data.persona); 
        if (recentPersonas.length > 5) { 
            recentPersonas = recentPersonas.slice(0, 5);
        }
        
        displayRecentPersonas(recentPersonas);
        
        fetchGenderCount();
        
        return data;
    } catch (error) {
        console.error('Error:', error);
        alert('Error al registrar persona: ' + error.message);
    }
}

function displayRecentPersonas(personas) {
    recentPersonasContainer.innerHTML = '';
    
    if (personas.length === 0) {
        recentPersonasContainer.innerHTML = '<p>No hay personas registradas recientemente</p>';
        return;
    }
    
    personas.forEach(persona => {
        const personaElement = document.createElement('div');
        personaElement.className = 'persona-item';
        personaElement.innerHTML = `
            <p><strong>Nombre:</strong> ${persona.nombre}</p>
            <p><strong>Edad:</strong> ${persona.edad}</p>
            <p><strong>Género:</strong> ${persona.genero}</p>
            <p><strong>Sexo:</strong> ${persona.sexo ? 'Hombre' : 'Mujer'}</p>
        `;
        recentPersonasContainer.appendChild(personaElement);
    });
}

async function fetchGenderCount() {
    try {
        const response = await fetch(`${API_BASE_URL}/personas/gender_count`);
        
        if (!response.ok) {
            throw new Error('Error al obtener contador de géneros');
        }
        
        const counts = await response.json();
        updateGenderCounts(counts);
    } catch (error) {
        console.error('Error:', error);
    }
}

function updateGenderCounts(counts) {
    maleCountElement.textContent = counts.hombres;
    femaleCountElement.textContent = counts.mujeres;
}

personaForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const personaData = {
        nombre: document.getElementById('nombre').value,
        edad: parseInt(document.getElementById('edad').value),
        sexo: document.getElementById('sexo').value === 'true',
        genero: document.getElementById('genero').value
    };
    
    await registerPersona(personaData);
    
    personaForm.reset();
});

+async function init() {
    recentPersonas = [];
    displayRecentPersonas(recentPersonas);
    
    await fetchGenderCount();
    
    setInterval(fetchGenderCount, 10000);
}

init();