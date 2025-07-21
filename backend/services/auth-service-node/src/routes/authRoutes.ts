import { Router } from 'express';
import { register, login } from '../controllers/authController';
import { protect } from '../middleware/authMiddleware';

const router = Router();

router.post('/register', register);
router.post('/login', login);

// Example of a protected route
router.get('/profile', protect, (req: any, res) => {
  res.status(200).json({ message: 'Welcome to your profile!', user: req.user });
});

export default router;