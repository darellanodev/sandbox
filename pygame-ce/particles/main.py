import pygame
import sys
from settings import *
from emitter import Emitter

def main():
    # Initialize Pygame
    pygame.init()
    screen = pygame.display.set_mode((SCREEN_WIDTH, SCREEN_HEIGHT))
    pygame.display.set_caption(TITLE)
    clock = pygame.time.Clock()

    # Create Emitter
    emitter = Emitter(SCREEN_WIDTH // 2, SCREEN_HEIGHT // 2)

    running = True
    while running:
        # 1. Event Handling
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False
            
            emitter.handle_input(event)

        # 2. Update
        dt = clock.tick(FPS) / 1000.0  # Delta time in seconds
        emitter.update(dt)

        # 3. Draw
        screen.fill(BG_COLOR)
        emitter.draw(screen)
        
        # Debug info
        debug_surf = pygame.font.SysFont(None, 24).render(f"Particles: {len(emitter.particles)}", True, (255, 255, 255))
        screen.blit(debug_surf, (10, 10))

        pygame.display.flip()

    pygame.quit()
    sys.exit()

if __name__ == "__main__":
    main()
