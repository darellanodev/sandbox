import pygame
import random
from particle import Particle
from settings import *

class Emitter:
    def __init__(self, x, y):
        self.pos = pygame.math.Vector2(x, y)
        self.particles = []
        self.emit_timer = 0
        self.emitting = True

    def emit(self):
        if self.emitting:
            # Random velocity
            vx = random.uniform(-1, 1)
            vy = random.uniform(-1, 1)
            vel = pygame.math.Vector2(vx, vy).normalize() * PARTICLE_SPEED
            
            # Random color variation (optional, keeping it white-ish)
            color = PARTICLE_COLOR
            
            self.particles.append(Particle(self.pos, vel, color, PARTICLE_LIFETIME))

    def update(self, dt):
        self.emit_timer += dt
        if self.emit_timer >= EMISSION_RATE:
            self.emit_timer = 0
            self.emit()
            
        # Update all particles
        for particle in self.particles:
            particle.update(dt)
            
        # Remove dead particles
        self.particles = [p for p in self.particles if p.alive]

    def draw(self, surface):
        for particle in self.particles:
            particle.draw(surface)
            
    def handle_input(self, event):
        if event.type == pygame.MOUSEMOTION:
            self.pos = pygame.math.Vector2(event.pos)
        elif event.type == pygame.MOUSEBUTTONDOWN:
            self.emitting = not self.emitting
