import pygame
import random
from settings import *

class Particle:
    def __init__(self, pos, vel, color, lifetime):
        self.pos = pygame.math.Vector2(pos)
        self.vel = pygame.math.Vector2(vel)
        self.color = color
        self.lifetime = lifetime
        self.start_lifetime = lifetime
        self.alive = True

    def update(self, dt):
        self.lifetime -= dt
        if self.lifetime <= 0:
            self.alive = False
        
        # Move particle
        self.pos += self.vel * dt

    def draw(self, surface):
        if self.alive:
            # Calculate alpha based on lifetime
            alpha = int(255 * (self.lifetime / self.start_lifetime))
            
            # Create a surface for the particle to support alpha
            particle_surf = pygame.Surface((PARTICLE_SIZE * 2, PARTICLE_SIZE * 2), pygame.SRCALPHA)
            pygame.draw.circle(particle_surf, (*self.color, alpha), (PARTICLE_SIZE, PARTICLE_SIZE), PARTICLE_SIZE)
            
            surface.blit(particle_surf, self.pos - (PARTICLE_SIZE, PARTICLE_SIZE))
