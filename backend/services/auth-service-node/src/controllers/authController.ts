import { Request, Response } from 'express';
import * as authService from '../services/authService';

export const register = async (req: Request, res: Response) => {
  try {
    const newUser = await authService.registerUser(req.body);
    // Exclude password hash from the response
    const { passwordHash, ...userResponse } = newUser;
    res.status(201).json(userResponse);
  } catch (error: any) {
    res.status(400).json({ message: error.message });
  }
};

export const login = async (req: Request, res: Response) => {
  try {
    const token = await authService.loginUser(req.body);
    if (!token) {
      return res.status(401).json({ message: 'Invalid credentials' });
    }
    res.status(200).json({ token });
  } catch (error: any) {
    res.status(500).json({ message: 'Server error' });
  }
};